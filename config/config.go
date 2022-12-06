package config

import (
	"errors"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Category string

const (
	A Category = "A"
	B Category = "B"
	C Category = "C"
	D Category = "D"
)

// Record defines a single record in the csv file
// with each field representing a single column
type Record struct {
	Name        string   `csv:"Name"`
	Phone       string   `csv:"Phone"`
	Email       string   `csv:"Email"`
	Notes       string   `csv:"Notes"`
	LastContact string   `csv:"Last Contact"`
	NextContact string   `csv:"Next Contact"`
	Category    Category `csv:"Category"`
}

// ReadCSV opens and reads a csv file and returns
// a list of the Record objects
func ReadCSV(path string) ([]*Record, error) {
	var file *os.File

	// Check if file exists
	_, err := os.Stat(path)

	// If file does not exist create it
	if errors.Is(err, os.ErrNotExist) {
		log.Println("file at given path does not exist, creating one")

		// Create the file at the given path
		file, err = os.Create(path)
		if err != nil {
			log.Println("failed to create file at given path", err)
			return nil, err
		}
	}

	// records will hold all the Record objects in the csv
	records := []*Record{}

	// Unamarshal csv into the records list
	// TODO: If file is empty, write the header row at the very top
	//       before reading from it and do this using error handling
	err = gocsv.UnmarshalFile(file, &records)
	if err != nil {
		log.Println("failed to unamarshal contents of the csv", err)
		return nil, err
	}

	return records, nil
}
