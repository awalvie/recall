package main

import (
	"log"

	"github.com/awalvie/recall/config"
)

func main() {
	// Path to csv file
	path := "data.csv"

	// Read csv file with records
	_, err := config.ReadCSV(path)
	if err != nil {
		log.Fatalf("faled to read file: %s", path)
	}

	log.Println("successfully read csv file", path)
}
