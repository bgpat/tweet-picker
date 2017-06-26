package main

import (
	"log"

	"github.com/bgpat/tweet-picker/client"
	"github.com/bgpat/twtr"
)

func main() {
	client, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Open(); err != nil {
		log.Fatal(err)
	}
	client.DeletedTweet = make(chan *twtr.Tweet)
	client.StreamingError = make(chan error)
	for {
		select {
		case tweet := <-client.DeletedTweet:
			log.Printf("tweet from @%s: %s\n", tweet.User.ScreenName, tweet.Text)
		case err := <-client.StreamingError:
			log.Printf("error: %+v\n", err)
		}
	}
}
