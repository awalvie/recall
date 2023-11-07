package mail

import (
	"log"
	"time"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/models"
	"github.com/emersion/go-smtp"
)

type Server struct {
	Host     string
	Port     int
	Username string
	Password string
	TLS      bool
}

func (s *Server) Start(a *config.App) {
	log.Println("starting mail server")

	// Check connection to the configured mail server
	_, err := smtp.Dial(s.Host)
	if err != nil {
		log.Println("error connecting to mail server")
	}

	// Get database connetion from app context
	db := a.DB

	// Get all contacts with next contact date as today
	var contacts []models.Contact
	today := time.Now().Format("2006-01-02")
	db.Where("DATE(next_contact) = ?", today).Find(&contacts)

	// Generate email from these contacts

}
