package database

import (
	"os"

	"github.com/go-redis/redis"
)

type Database struct {
	*redis.Client
	URL string
}

func New() (*Database, error) {
	db := Database{
		URL: os.Getenv("REDIS_URL"),
	}
	option, err := redis.ParseURL(db.URL)
	if err != nil {
		return nil, err
	}
	db.Client = redis.NewClient(option)
	return &db, nil
}

func (db *Database) GetString(key, defaultValue string) string {
	if value, err := db.Get(key).Result(); err == nil {
		return value
	}
	return defaultValue
}
