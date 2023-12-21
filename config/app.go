package config

import (
	"time"

	"gorm.io/gorm"
)

// App holds the application configuration and is
// passed to the handlers
type App struct {
	Config   *Config
	DB       *gorm.DB
	NextMail time.Time
}
