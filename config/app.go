package config

import "gorm.io/gorm"

type App struct {
	Config *Config
	DB     *gorm.DB
}
