package gollection

import (
	"fmt"
	"gopkg.in/redis.v3"
	"log"
)

func NewRedis(c Config) (*redis.Client, error) {
	opt := &redis.Options{
		Addr: fmt.Sprintf("%s:%d", c.RedisConfig.Host, c.RedisConfig.Port),
		DB:   int64(c.RedisConfig.Database),
	}

	if c.RedisConfig.Password != "" {
		log.Println("%v", c.RedisConfig)
		opt.Password = c.RedisConfig.Password
	}

	return redis.NewClient(opt), nil
}
