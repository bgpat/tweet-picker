package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bgpat/tweet-picker/client"
	"github.com/bgpat/twtr"
)

const (
	expiration = time.Hour * 24
)

func main() {
	client, err := client.New()
	if err != nil {
		log.Fatal(err)
	}
	for {
		streaming, err := client.Streaming()
		defer streaming.Close()
		if err != nil {
			log.Fatal(err)
		}
		for {
			event, err := streaming.Decode()
			if err != nil {
				log.Println("error: ", err.Error())
				break
			}
			switch data := event.(type) {
			case *twtr.Tweet:
				id := data.IDStr
				tweet := fmt.Sprintf("@%s %s", data.User.ScreenName, data.Text)
				err := client.Database.Set(id, tweet, expiration).Err()
				if err != nil {
					log.Println("error: ", err.Error())
					break
				}
			case *twtr.StreamingTweetEvent:
				id := data.TargetObject.IDStr
				tweet := fmt.Sprintf("@%s %s", data.TargetObject.User.ScreenName, data.TargetObject.Text)
				err := client.Database.Set(id, tweet, expiration).Err()
				if err != nil {
					log.Println("error: ", err.Error())
					break
				}
			case *twtr.StreamingDeleteTweetEvent:
				tweet := client.Database.GetString(data.Delete.Status.IDStr, data.Delete.Status.IDStr)
				log.Printf("delete tweet: %s\n", tweet)
			default:
				log.Printf("continue: %T\n", event)
			}
		}
	}
}
