package client

import (
	"os"

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
}

func New() (*Client, error) {
	c, err := cache.New()
	if err != nil {
		return nil, err
	}
	consumer := twtr.NewCredentials(consumerKey, consumerSecret)
	token := twtr.NewCredentials(accessToken, accessSecret)
	client := Client{
		Cache:  c,
		Client: twtr.NewClient(consumer, token),
	}
	return &client, nil
}

func (c *Client) Streaming() (*twtr.Streaming, error) {
	params := twtr.Params{
		"stall_warnings": "true",
		"filter_level":   "none",
		"replies":        "all",
	}
	return c.NewUserStream(&params)
}
