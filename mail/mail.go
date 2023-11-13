package mail

import (
	"bytes"
	"log"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/models"
	"github.com/emersion/go-sasl"
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
	for {
		// Run once a day
		time.Sleep(24 * time.Hour)

		log.Println("starting mail server")

		// Get contacts that need to be contacted today
		contacts, err := models.GetContactsToday(a.DB)
		if err != nil {
			log.Println("error getting contacts to be contacted today", err)
		}

		// Generate email from template
		// Get the template path
		tpath := filepath.Join(a.Config.Dirs.Templates, "*")

		// Parse the teamplates
		t := template.Must(template.ParseGlob(tpath))

		// Generate struct to pass to the template
		data := struct {
			To       []string
			From     string
			Subject  string
			Contacts []models.Contact
		}{
			To:       a.Config.Mail.To,
			From:     a.Config.Mail.From,
			Subject:  a.Config.Mail.Subject,
			Contacts: contacts,
		}

		// Execute the template
		var emailBody bytes.Buffer
		err = t.ExecuteTemplate(&emailBody, "email", data)
		if err != nil {
			log.Println("error executing template", err)
		}

		// Remove the leading new line
		emailBodyTrim := bytes.TrimLeft(emailBody.Bytes(), "\n")

		// Server address
		addr := s.Host + ":" + strconv.Itoa(s.Port)

		// Authenticate to the mail Server
		auth := sasl.NewPlainClient("", s.Username, s.Password)

		// Send email
		err = smtp.SendMail(addr, auth, a.Config.Mail.From, a.Config.Mail.To, bytes.NewReader(emailBodyTrim))
		if err != nil {
			log.Println("error sending mail", err)
		} else {
			log.Println("mail sent successfully")
		}
	}
}
