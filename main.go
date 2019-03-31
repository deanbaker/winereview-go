package main

import (
	"log"
	"os"

	"github.com/deandemo/winereview/wine/http"

	"github.com/deandemo/winereview/wine/db"

	"github.com/deandemo/winereview/wine/io"
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

	// Swapping out the Reviewer Interface
	reviewer := db.NewSQLCache()
	// reviewer := db.NewMemStore()

	io.Parse(file, reviewer)

	http.Serve(reviewer)
}
