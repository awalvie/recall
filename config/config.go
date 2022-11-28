package config

import (
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

type Person struct {
	Name        string   `csv:"name"`
	Phone       string   `csv:"phone"`
	Email       string   `csv:"email"`
	Notes       string   `csv:"notes"`
	LastContact string   `csv:"last_contact"`
	NextContact string   `csv:"next_contact"`
	Category    Category `csv:"category"`
}

// ReadCSV opens and reads a csv files and returns
// the output as a list of records
func ReadCSV(path string) ([]*Person, error) {
	// Open the file on the path
	// TODO: If file is just created populate the first
	// 	     line with column titles
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln("failed to open file", path, err)
	}
	defer file.Close()

	// people will hold all the resocrds in the csv
	people := []*Person{}

	// Unamarshal csv into people struct
	err = gocsv.UnmarshalFile(file, &people)
	if err != nil {
		log.Fatalln("failed to unmarshal csv file", path, err)
	}

	return people, nil
}
