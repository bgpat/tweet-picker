package models

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type Tweet struct {
	ID        int64 `gorm:"primary_key,AUTO_INCREMENT"`
	UserID    int64
	User      *User
	CreatedAt *time.Time
	DeletedAt time.Time
	JSON      string
}

func (t *Tweet) MarshalJSON() ([]byte, error) {
	if t.JSON == "" {
		return json.Marshal(map[string]interface{}{
			"id":          t.ID,
			"id_str":      strconv.FormatInt(t.ID, 10),
			"user_id":     t.UserID,
			"user_id_str": strconv.FormatInt(t.UserID, 10),
			"deleted_at":  t.DeletedAt,
		})
	}
	tweet := map[string]interface{}{}
	err := json.Unmarshal([]byte(t.JSON), &tweet)
	if err != nil {
		log.Fatalf("%+v\n%+v\n", err, t.JSON)
		return nil, err
	}
	tweet["deleted_at"] = t.DeletedAt
	return json.Marshal(tweet)
}
