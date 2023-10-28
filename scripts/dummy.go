// Populate the database with dummy data.

package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/awalvie/recall/models"
	"github.com/go-faker/faker/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var dbFile string

	// Parse flags
	flag.StringVar(&dbFile, "database", "./sqlite.db", "path to database file")
	flag.Parse()

	flag.Usage = func() {
		log.Println("Usage: dummy [options]")
		flag.PrintDefaults()
	}

	// Connect to database
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	}
	log.Println("connected to database:", dbFile)

	// Delete existing tables
	err = db.Migrator().DropTable(&models.Contact{})
	if err != nil {
		log.Fatalln("error dropping table:", err)
	}
	log.Println("dropped table:", "contacts")

	// Create tables
	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatalln("error creating table:", err)
	}
	log.Println("created table:", "contacts")

	// Generate contacts
	var contacts []models.Contact
	numContacts := 100
	for i := 0; i < numContacts; i++ {
		var contact models.Contact
		err := faker.FakeData(&contact)

		if err != nil {
			log.Fatalln("error generating contact:", err)
		}

		// Customize or generate specific fields as needed
		contact.LastContact = time.Now().AddDate(0, 0, -rand.Intn(365))
		contact.NextContact = time.Now().AddDate(0, 0, rand.Intn(365))
		contact.ID = uint(i + 1)

		contacts = append(contacts, contact)
	}
	log.Println("generated", numContacts, "contacts")

	// Write generated data to database
	err = db.Create(&contacts).Error

}
