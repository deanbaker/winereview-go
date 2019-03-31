package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/deandemo/winereview/wine"
)

func main() {

	log.Println("Starting application")
	defer log.Println("Finished the application")

	// Open the file, defer the close
	file, err := os.Open("resources/winemag-data-20-v2.csv")
	defer file.Close()
	if err != nil {
		log.Fatal("Could not read file")
	}

	// Create a csv reader from the given file
	r := csv.NewReader(file)
	_, err = r.Read()
	if err != nil {
		log.Fatal("Could not create the csv reader")
	}

	// Read all records of the csv
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("Could not read")
	}

	// Create a structure to hold the top candidates
	topReviewer := make(map[string]int)
	topVariety := make(map[string]int)

	// Note that we aren't using []wine.Review here - we can use the wine.Reviews structure
	rr := make(wine.Reviews, 0)
	for _, line := range records {
		points, err := strconv.Atoi(line[4])
		// Default points to 0 if nott present
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

		// Increment our counters
		topReviewer[review.TasterName]++
		topVariety[review.Variety]++
		rr = append(rr, review)
	}

	log.Printf("Number of reviews: %d", len(rr))
	log.Printf("Most prolific Reviewer: %s", top(topReviewer))
	log.Printf("Most reviewed variety: %s", top(topVariety))
}

// Iterate through the map and select the highest candidate
// Note that maps do not enforce order - useage may vary
// Note that we are using a naked return statement - https://tour.golang.org/basics/7
func top(m map[string]int) (top string) {

	highest := 0
	for k, v := range m {
		if v > highest {
			top = k
			highest = v
		}
	}
	return
}
