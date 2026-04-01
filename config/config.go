package config

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Config struct {
	db    *gorm.DB
	cache *redis.Client
}
