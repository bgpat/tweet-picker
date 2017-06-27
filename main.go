package main

import (
	"encoding/gob"
	"log"

	"github.com/bgpat/tweet-picker/client"
	"github.com/bgpat/tweet-picker/database"
	"github.com/bgpat/tweet-picker/models"
	"github.com/bgpat/tweet-picker/server"
)

func main() {
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
	if _, err := database.New(); err != nil {
		log.Fatalf("exit process: %s", err.Error())
	}
	go func() {
		err := startServer()
		log.Fatalf("exit process: %s", err.Error())
	}()
	err := startClient()
	log.Fatalf("exit process: %s", err.Error())
}

func startClient() error {
	c, err := client.New()
	if err != nil {
		return err
	}
	if err := c.Open(); err != nil {
		return err
	}
	c.DeletedTweet = make(chan *client.Tweet)
	c.StreamingError = make(chan error)
	db := database.Default()
	for {
		select {
		case tweet := <-c.DeletedTweet:
			model, err := tweet.Model()
			if err != nil {
				log.Printf("error: %+v\n", err)
				continue
			}
			db.Create(model)
			if tweet.Tweet == nil {
				log.Printf("tweet from %d\n", tweet.UserID)
				continue
			}
			user := models.User{}
			userModel, err := tweet.UserModel()
			if err != nil {
				log.Printf("error: %+v\n", err)
				continue
			}
			db.Where(models.User{
				ID: tweet.UserID,
			}).Assign(userModel).FirstOrCreate(&user)
			db.Save(&user)
			log.Printf("tweet from @%s: %s\n", tweet.User.ScreenName, tweet.Text)
		case err := <-c.StreamingError:
			log.Printf("error: %+v\n", err)
		}
	}
	return nil
}

func startServer() error {
	s := server.New()
	return s.Run()
}
