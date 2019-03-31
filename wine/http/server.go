package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/deandemo/winereview/wine"
)

// Serve will map out CRUD routes for the Review resource and start a http server.
func Serve(reviewer wine.Reviewer) {
	log.Println("Starting http server")
	http.HandleFunc("/reviews", func(w http.ResponseWriter, r *http.Request) {
		// Allow CORS for swagger testing
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		// w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(reviewer.FindAll())
		case "POST":
			log.Println("Creating")

			review := wine.Review{}
			err := json.NewDecoder(r.Body).Decode(&review)

			if err != nil {
				http.Error(w, "Bad Request", 400)
				return
			}

			review, err = reviewer.Save(review)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			json.NewEncoder(w).Encode(review)
		case "OPTIONS":
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			fmt.Fprintln(w, "ok")
		default:
			http.Error(w, "Not supported", 405)
			return
		}
	})

	http.HandleFunc("/reviews/", func(w http.ResponseWriter, r *http.Request) {

		// Allow CORS for swagger testing
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		// w.Header().Set("Content-Type", "application/json")
		param := strings.TrimPrefix(r.URL.Path, "/reviews/")

		id, err := strconv.Atoi(param)
		if err != nil {
			log.Printf("Could not parse path [%s]", param)
			return
		}

		log.Printf("id [%d]", id)
		review, ok := reviewer.Find(id)
		if !ok {
			log.Println("Could not find review")
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case "GET":
			json.NewEncoder(w).Encode(review)

		case "PUT":
			log.Printf("PUT")
			err := json.NewDecoder(r.Body).Decode(&review)
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			_, ok := reviewer.Update(review)
			if !ok {
				http.Error(w, err.Error(), 400)
				return
			}
			json.NewEncoder(w).Encode(review)

		case "DELETE":
			log.Printf("DELETE")

			ok := reviewer.Delete(id)
			if !ok {
				http.Error(w, "Could not delete", 400)
				return
			}

			fmt.Fprintln(w, "OK")
			return

		case "OPTIONS":
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			fmt.Fprintln(w, "ok")

		default:
			http.Error(w, "Not supported", 405)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
