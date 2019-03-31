package io

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/deandemo/winereview/wine"
)

// Parse will take in a Reader and save its contents into the Saver.
func Parse(input io.Reader, db wine.Saver) {

	log.Println("Starting the parser")

	// Create a csv reader from the given file
	r := csv.NewReader(input)
	_, err := r.Read()
	if err != nil {
		log.Fatal("Could not create the csv reader")
	}

	// Read all records of the csv
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Could not read")
	}

	for _, line := range records {
		// Default points to zero
		points, err := strconv.Atoi(line[4])
		if err != nil {
			points = 0
		}
		review := wine.Review{
			Title:      line[11],
			Country:    line[1],
			Variety:    line[12],
			Points:     points,
			TasterName: line[9],
		}
		db.Save(review)
	}
}
