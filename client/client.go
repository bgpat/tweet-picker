package client

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/bgpat/tweet-picker/client/cache"
	"github.com/bgpat/twtr"
)

var (
	consumerKey    = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken    = os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret   = os.Getenv("TWITTER_ACCESS_SECRET")
)

type Client struct {
	*cache.Cache
	*twtr.Client
	*twtr.Streaming
	Expiration     time.Duration
	DeletedTweet   chan *Tweet
	StreamingError chan error
}

func New() (*Client, error) {
	c, err := cache.New()
	if err != nil {
		return nil, err
	}
	consumer := twtr.NewCredentials(consumerKey, consumerSecret)
	token := twtr.NewCredentials(accessToken, accessSecret)
	expiration, err := strconv.ParseInt(os.Getenv("CACHE_EXPIRATION"), 10, 64)
	if err != nil {
		return nil, err
	}
	client := Client{
		Cache:      c,
		Client:     twtr.NewClient(consumer, token),
		Expiration: time.Duration(expiration) * time.Second,
	}
	return &client, nil
}

func (c *Client) Open() error {
	if c.Streaming != nil {
		return errors.New("streaming has opened already")
	}
	s, err := c.NewUserStream(&twtr.Params{
		"stall_warnings": "true",
		"filter_level":   "none",
		"replies":        "all",
	})
	if err != nil {
		return err
	}
	c.Streaming = s
	go c.decodeStreaming()
	return nil
}

func (c *Client) decodeStreaming() {
	for {
		event, err := c.Streaming.Decode()
		if err == nil {
			switch data := event.(type) {
			case *twtr.Tweet:
				err = c.storeTweet(data)
			case *twtr.StreamingTweetEvent:
				err = c.storeTweet(&data.TargetObject)
			case *twtr.StreamingDeleteTweetEvent:
				err = c.pickupTweet(data)
			default:
				err = fmt.Errorf("unhandled event: %+v", data)
			}
		}
		if err != nil {
			c.StreamingError <- err
		}
	}
}

func (c *Client) storeTweet(tweet *twtr.Tweet) error {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(tweet)
	if err != nil {
		return err
	}
	c.Cache.Set(tweet.IDStr, buf.String(), c.Expiration)
	return nil
}

func (c *Client) pickupTweet(event *twtr.StreamingDeleteTweetEvent) error {
	timestamp, err := strconv.ParseInt(event.Delete.TimestampMS, 10, 64)
	deletedAt := time.Unix(timestamp/1000, (timestamp-timestamp/1000*1000)*10e6)
	buf, err := c.Cache.Get(event.Delete.Status.IDStr).Bytes()
	if err != nil {
		c.DeletedTweet <- &Tweet{
			ID:        event.Delete.Status.ID.ID,
			UserID:    event.Delete.Status.UserID,
			DeletedAt: deletedAt,
		}
		return err
	}
	tweet := twtr.Tweet{}
	decoder := gob.NewDecoder(bytes.NewBuffer(buf))
	if err := decoder.Decode(&tweet); err != nil {
		return err
	}
	c.DeletedTweet <- &Tweet{
		Tweet:     &tweet,
		ID:        event.Delete.Status.ID.ID,
		UserID:    event.Delete.Status.UserID,
		DeletedAt: deletedAt,
	}
	return nil
}
