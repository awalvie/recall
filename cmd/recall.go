package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/mail"
	"github.com/awalvie/recall/models"
	"github.com/awalvie/recall/routes"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// init runs before main
func init() {
	// Configure logger flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// main is the entrypoint for the application
func main() {
	var cfg, dbFile string
	var err error

	// Configure CLI flags
	flag.StringVar(&cfg, "config", "./config.yaml", "path to config file")
	flag.StringVar(&dbFile, "database", "./sqlite.db", "path to database file")
	flag.Parse()

	// Read config file
	c, err := config.Read(cfg)
	if err != nil {
		log.Fatalln("error reading config file:", err)
	}
	log.Println("config file read successfully")

	// Connect to database
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	}

	// Migrate database
	if err := db.AutoMigrate(
		&models.Contact{},
	); err != nil {
		log.Fatalln("error migrating database:", err)
	}

	// Create app config
	app := config.App{
		Config: c,
		DB:     db,
	}

	// Initialize mail server
	mailServer := mail.Server{
		Host:     c.Mail.Host,
		Port:     c.Mail.Port,
		Username: c.Mail.Username,
		Password: c.Mail.Password,
		TLS:      c.Mail.TLS,
	}

	go mailServer.Start(&app)

	// Configure HTTP server
	e := echo.New()

	// Configure echo server
	routes.Configure(e, &app)
	log.Println("server configured successfully")

	// Configure server address
	addr := fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)

	// Start HTTP server
	log.Println("starting server on", addr)

	err = e.Start(addr)
	if err != nil {
		log.Fatalln("error starting server:", err)
	}
}
