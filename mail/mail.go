package mail

import (
	"bytes"
	"log"
	"path/filepath"
	"strconv"
	"text/template"

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
	log.Println("starting mail server")

	addr := s.Host + ":" + strconv.Itoa(s.Port)

	// Check connection to the configured mail server
	client, err := smtp.Dial(addr)
	if err != nil {
		log.Println("error connecting to mail server", err)
	}
	defer client.Close()
	log.Println("connected to mail server")

	// Authenticate to the mail Server
	auth := sasl.NewPlainClient("", s.Username, s.Password)

	// Get contacts that need to be contacted today
	contacts, err := models.GetContactsToday(a.DB)
	if err != nil {
		log.Println("error getting contacts to be contacted today", err)
	}
	log.Println("contacts to be contacted today", contacts)

	// Generate email from template
	// Get the template path
	tpath := filepath.Join(a.Config.Dirs.Templates, "*")

	// Parse the teamplates
	t := template.Must(template.ParseGlob(tpath))

	// Generate struct to pass to the template
	data := struct {
		To      []string
		From    string
		Subject string
	}{
		To:      a.Config.Mail.To,
		From:    a.Config.Mail.From,
		Subject: "Damn, go templates are awesome!",
	}

	// Execute the template
	var emailBody bytes.Buffer
	err = t.ExecuteTemplate(&emailBody, "email", data)
	if err != nil {
		log.Println("error executing template", err)
	}

	// Send email
	err = smtp.SendMail(addr, auth, a.Config.Mail.From, a.Config.Mail.To, &emailBody)
	if err != nil {
		log.Println("error sending mail", err)
	}
	log.Println("email sent")

}
