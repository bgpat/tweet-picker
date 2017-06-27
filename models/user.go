package models

import (
	"encoding/json"
	"strconv"
)

type User struct {
	ID         int64    `gorm:"primary_key,AUTO_INCREMENT" json:"id,string"`
	ScreenName string   `json:"screen_name"`
	JSON       string   `json:"json"`
	Tweets     []*Tweet `json:"tweets,omitempty"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	if u.JSON == "" {
		return json.Marshal(map[string]interface{}{
			"id":     u.ID,
			"id_str": strconv.FormatInt(u.ID, 10),
		})
	}
	return []byte(u.JSON), nil
}
