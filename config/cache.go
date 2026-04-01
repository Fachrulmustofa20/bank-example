package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func (cfg *Config) initCache() error {
	redis := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Username: "default",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := redis.Ping(context.Background()).Err()
	if err != nil {
		return err
	}

	cfg.cache = redis
	log.Println("Redis connected successfully...")
	return nil
}
