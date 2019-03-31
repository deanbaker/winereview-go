package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/deandemo/winereview/wine"
	// Black import to load the sqlite3 driver
	// https://golang.org/doc/effective_go.html#init
	_ "github.com/mattn/go-sqlite3"
)

// SQLCache will implement all the things
type SQLCache struct {
	db *sql.DB
}

// FindAll will return all elements in the database
func (c *SQLCache) FindAll() (reviews wine.Reviews) {
	log.Println("Finding all by db")
	rows, err := c.db.Query("select id, title, variety, country, points, taster_name from reviews")
	defer rows.Close()

	reviews = constructResults(rows)
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return reviews
}

// Save will create a new record in the database
func (c *SQLCache) Save(r wine.Review) (review wine.Review, err error) {

	// Create insert statement and close
	stmt, err := c.db.PrepareContext(context.TODO(),
		`insert into 
		reviews (title, variety, country, points, taster_name) 
		values(?, ?, ?, ?, ?)`)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Exec(r.Title, r.Variety, r.Country, r.Points, r.TasterName)
	if err != nil {
		log.Println("Could not create Review")
		return
	}

	// Extract the last insert id
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("No id incremented")
		return wine.Review{}, errors.New("Did not create an id")
	}

	r.ID = int(id)
	return r, nil
}

// Delete will delete the review by id
func (c *SQLCache) Delete(i int) bool {

	log.Printf("Deleting review by [%d]", i)
	stmt, err := c.db.PrepareContext(context.TODO(), "delete from reviews WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(i)
	if err != nil {
		return false
	}

	return true
}

// Update will update the given Review in the database
func (c *SQLCache) Update(r wine.Review) (review wine.Review, ok bool) {

	stmt, err := c.db.PrepareContext(context.TODO(),
		`UPDATE reviews 
		SET title = ?, variety = ?, country = ?, points = ?, taster_name = ? 
		WHERE id = ?`)
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(r.Title, r.Variety, r.Country, r.Points, r.TasterName, r.ID)
	if err != nil {
		log.Println("Could not update")
		return
	}

	return review, true
}

// Find will return a Review from the database given its ID. It will return false
// if none exists or if there is an error
func (c *SQLCache) Find(i int) (wine.Review, bool) {

	log.Printf("Find review by [%d]", i)

	// Statement should be prepared at a higher level as it is reusable
	stmt, err := c.db.PrepareContext(context.TODO(), "select id, title, variety, country, points, taster_name from reviews WHERE id = ?")
	defer stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	r, err := scan(stmt.QueryRow(i))
	if err != nil {
		log.Println("Could not find the thing", err)
		return wine.Review{}, false
	}
	return r, true
}

// NewSQLCache will connect to the database and return.
func NewSQLCache() *SQLCache {

	return &SQLCache{db: connect()}
}

// Scanner interface so we can use both sql.Row and sql.Rows
type Scanner interface {
	Scan(dest ...interface{}) error
}

// scan will create a new Review object from a given resultset
func scan(s Scanner) (review wine.Review, err error) {

	// Note that we are passing in pointers to the data we want to be set on
	// out destination object review
	err = s.Scan(&review.ID,
		&review.Title,
		&review.Variety,
		&review.Country,
		&review.Points,
		&review.TasterName)

	return
}

func constructResults(rows *sql.Rows) wine.Reviews {

	results := make(wine.Reviews, 0)
	for rows.Next() {
		review, err := scan(rows)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, review)
	}

	return results
}

// Connect will remove any old sqllite files, and create a new connection.
func connect() *sql.DB {

	os.Remove("./reviews.db")
	db, err := sql.Open("sqlite3", "./reviews.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println("Could not connect to the database")
		log.Fatal(err)
	}

	initialise(db)
	return db
}

// initialise will create our reviews table.
func initialise(db *sql.DB) {

	sqlStmt := `
	create table reviews (
		id integer not null primary key, 
		title text, 
		variety text,
		country text, 
		points number, 
		taster_name text
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
