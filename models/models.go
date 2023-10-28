package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" faker:"oneof: 1,2",unique`
	FirstName   string `faker:"first_name"`
	LastName    string `faker:"last_name"`
	Email       string `faker:"email"`
	Phone       string `faker:"phone_number"`
	Category    string `faker:"oneof: A, B, C, D"`
	LastContact time.Time
	NextContact time.Time
	Notes       string `faker:"paragraph"`
}
