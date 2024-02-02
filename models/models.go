package models

import (
	"errors"
	"net/mail"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" faker:"oneof: 1,2",unique`
	Name        string `faker:"name" form:"name" binding:"required"`
	Email       string `faker:"email" form:"email" binding:"required,email"`
	Phone       string `faker:"phone_number" form:"phone"`
	Category    string `faker:"oneof: A, B, C, D" form:"category"`
	LastContact time.Time
	NextContact time.Time
	Notes       string `faker:"paragraph" form:"notes"`
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (c *Contact) Validate() error {
	if c.Email != "" {
		// Check if email is valid
		_, err := mail.ParseAddress(c.Email)
		if err != nil {
			return err
		}

	}
	// Check if phone is valid
	if c.Phone != "" {
		re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
		if !re.MatchString(c.Phone) {
			return errors.New("Invalid phone number")
		}

	}
	return nil
}

func (c *Contact) GetNextContactDate() time.Time {
	var nextContact time.Time
	if c.LastContact.IsZero() {
		switch c.Category {
		case "A":
			nextContact = time.Now().AddDate(0, 0, 7)
		case "B":
			nextContact = time.Now().AddDate(0, 0, 21)
		case "C":
			nextContact = time.Now().AddDate(0, 0, 30)
		case "D":
			nextContact = time.Now().AddDate(0, 0, 90)
		}
	} else {
		switch c.Category {
		case "A":
			nextContact = c.LastContact.AddDate(0, 0, 7)
		case "B":
			nextContact = c.LastContact.AddDate(0, 0, 21)
		case "C":
			nextContact = c.LastContact.AddDate(0, 0, 30)
		case "D":
			nextContact = c.LastContact.AddDate(0, 0, 90)
		}
	}
	return nextContact
}

func GetContactsToday(db *gorm.DB) ([]Contact, error) {
	// Get all contacts with next contact date as today
	var contacts []Contact
	today := time.Now().Format("2006-01-02")
	db.Where("DATE(next_contact) = ?", today).Find(&contacts)
	return contacts, nil
}

// GetUpcomingContacts returns all contacts with next contact date
// within the next 10 days
func GetUpcomingContacts(db *gorm.DB) ([]Contact, error) {
	// Get all contacts with next contact date as today
	var contacts []Contact
	today := time.Now().Format("2006-01-02")
	tenDays := time.Now().AddDate(0, 0, 10).Format("2006-01-02")
	db.Where("DATE(next_contact) BETWEEN ? AND ?", today, tenDays).Find(&contacts)
	return contacts, nil
}

// GetRecentContacts returns all contacts with last contact date
// within the last 10 days
func GetRecentContacts(db *gorm.DB) ([]Contact, error) {
	// Get all contacts with next contact date as today
	var contacts []Contact
	today := time.Now().Format("2006-01-02")
	tenDays := time.Now().AddDate(0, 0, -10).Format("2006-01-02")
	db.Where("DATE(last_contact) BETWEEN ? AND ?", today, tenDays).Find(&contacts)
	return contacts, nil
}
