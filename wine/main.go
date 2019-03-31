// Package wine holds all of the domain and interfaces
// that are used for subsequent packages.
//
// Note that this root package does not have any external dependencies
package wine

// Reviews is a slice of Review for convenience.
type Reviews []Review

// Review is the base domain object.
type Review struct {
	ID         int    `json:"_id,omitempty"`
	Title      string `json:"title,omitempty"`
	Variety    string `json:"variety,omitempty"`
	Country    string `json:"country,omitempty"`
	Points     int    `json:"points,omitempty"`
	TasterName string `json:"taster_name,omitempty"`
}

// Finder will find one or more reviews.
type Finder interface {
	FindAll() Reviews
	Find(int) (Review, bool)
}

// Saver will save a Review.
type Saver interface {
	Save(Review) (Review, error)
}

// Deleter will delete a review
type Deleter interface {
	Delete(int) bool
}

// Updater will update an interface
type Updater interface {
	Update(Review) (Review, bool)
}

// Reviewer composes all the things
type Reviewer interface {
	Finder
	Saver
	Deleter
	Updater
}
