package config

import "gorm.io/gorm"

// App holds the application configuration and is
// passed to the handlers
type App struct {
	Config *Config
	DB     *gorm.DB
}
