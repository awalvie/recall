package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/routes"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// Configure logger flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

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

	// Create app config
	app := config.App{
		Config: c,
		DB:     db,
	}

	// Configure HTTP server
	e := echo.New()

	// Configure echo server
	routes.Configure(e, &app)
	log.Println("server configured successfully")

	// Configure server address
	addr := fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)

	// Start HTTP server
	log.Println("starting server on", addr)
	e.Start(addr)
}
