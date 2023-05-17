package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/awalvie/recall/config"
	"github.com/awalvie/recall/routes"
	"github.com/labstack/echo/v4"
)

func init() {
	// Configure logger flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	var cfg string
	var err error

	// Configure CLI flags
	flag.StringVar(&cfg, "config", "./config.yaml", "path to config file")
	flag.Parse()

	// Read config file
	c, err := config.Read(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("config file read successfully")

	// Configure HTTP server
	e := echo.New()

	// Configure echo server
	routes.Configure(e)

	// Configure server address
	addr := fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)

	// Start HTTP server
	e.Start(addr)
}
