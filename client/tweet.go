package client

import (
	"encoding/json"
	"time"

	"github.com/bgpat/tweet-picker/models"
	"github.com/bgpat/twtr"
)

type Tweet struct {
	*twtr.Tweet
	ID        int64
	UserID    int64
	DeletedAt time.Time
}

func (t *Tweet) Model() (*models.Tweet, error) {
	tweet := &models.Tweet{
		ID:        t.ID,
		UserID:    t.UserID,
		DeletedAt: t.DeletedAt,
	}
	if t.Tweet != nil {
		buf, err := json.Marshal(t.Tweet)
		if err != nil {
			return nil, err
		}
		tweet.JSON = string(buf)
	}
	return tweet, nil
}

func (t *Tweet) UserModel() (*models.User, error) {
	user := &models.User{
		ID: t.UserID,
	}
	if t.Tweet != nil {
		user.ScreenName = t.Tweet.User.ScreenName
		buf, err := json.Marshal(t.Tweet.User)
		if err != nil {
			return nil, err
		}
		user.JSON = string(buf)
	}
	return user, nil
}
