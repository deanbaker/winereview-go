package db

import (
	"sync"

	"github.com/deandemo/winereview/wine"
)

// MemStore is an in memory data store for wine.Review objects
type MemStore struct {
	db           map[int]wine.Review
	currentIndex int
	mutex        sync.Mutex
}

// FindAll will iterate through the map and return all of the values.
// Note that order may not be guaranteed
func (s *MemStore) FindAll() wine.Reviews {

	rr := make([]wine.Review, 0)
	for _, r := range s.db {
		rr = append(rr, r)
	}
	return rr
}

// Find will find a Review by its associated ID.
func (s *MemStore) Find(i int) (wine.Review, bool) {

	r, ok := s.db[i]
	if !ok {
		return wine.Review{}, false
	}

	return r, true
}

// Save will generate an id and save the review in the map.
func (s *MemStore) Save(r wine.Review) (wine.Review, error) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.currentIndex++
	r.ID = s.currentIndex
	s.db[r.ID] = r

	return r, nil
}

// Delete will remove the element from the map.
func (s *MemStore) Delete(id int) bool {

	if _, ok := s.db[id]; !ok {
		return false
	}

	// GO's inbuilt delete function, because you cannot dereference a
	// concrete struct from a map
	delete(s.db, id)
	return true
}

// Update will replace the element in the map by its ID.
func (s *MemStore) Update(r wine.Review) (wine.Review, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	_, ok := s.db[r.ID]
	if !ok {
		return r, false
	}

	s.db[r.ID] = r
	return wine.Review{}, true
}

// NewMemStore will create a new empty MemStore
func NewMemStore() *MemStore {

	return &MemStore{db: make(map[int]wine.Review)}
}
