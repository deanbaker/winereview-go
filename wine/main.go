// Package wine holds all of the domain and interfaces
// that are used for subsequent packages.
//
// Note that this root package does not have any external dependencies
package wine

// Reviews is a slice of Review for convenience.
type Reviews []Review

// Review is the base domain object.
type Review struct {
	ID         int
	Title      string
	Variety    string
	Country    string
	Points     int
	TasterName string
}
