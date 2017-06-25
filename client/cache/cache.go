package cache

import (
	"os"

	"github.com/go-redis/redis"
)

type Cache struct {
	*redis.Client
	URL string
}

func New() (*Cache, error) {
	c := Cache{
		URL: os.Getenv("REDIS_URL"),
	}
	option, err := redis.ParseURL(c.URL)
	if err != nil {
		return nil, err
	}
	c.Client = redis.NewClient(option)
	return &c, nil
}

func (c *Cache) GetString(key, defaultValue string) string {
	if value, err := c.Get(key).Result(); err == nil {
		return value
	}
	return defaultValue
}
