package config

import (
	"gorm.io/gorm"
)

type Config struct {
	db *gorm.DB
}
