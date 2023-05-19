package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Category    string
	LastContact time.Time
	NextContact time.Time
	Notes       string
}
