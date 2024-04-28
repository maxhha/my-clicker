package storages

import (
	"fmt"

	"github.com/maxhha/my-clicker/internal/interactors/ports"
	"github.com/teris-io/shortid"
)

// Not thread safe
type InMemoryLinkStorage struct {
	links    map[string]string
	counters map[string]uint64
}

func NewInMemoryLinkStorage() InMemoryLinkStorage {
	return InMemoryLinkStorage{
		links:    make(map[string]string),
		counters: make(map[string]uint64),
	}
}

// Create implements ports.LinkStorage.
func (s *InMemoryLinkStorage) Create(redirect string) (string, error) {
	id, err := shortid.Generate()
	if err != nil {
		return "", fmt.Errorf("shortid generate: %w", err)
	}

	s.links[id] = redirect
	s.counters[id] = 0

	return id, nil
}

// GetCounter implements ports.LinkStorage.
func (s *InMemoryLinkStorage) GetCounter(link string) (uint64, error) {
	counter, ok := s.counters[link]
	if ok {
		return counter, nil
	} else {
		return 0, ports.ErrNotExists
	}
}

// GetRedirect implements ports.LinkStorage.
func (s *InMemoryLinkStorage) GetRedirect(link string) (string, error) {
	redirect, ok := s.links[link]
	if ok {
		return redirect, nil
	} else {
		return "", ports.ErrNotExists
	}
}

// IncrementCounter implements ports.LinkStorage.
func (s *InMemoryLinkStorage) IncrementCounter(link string) error {
	counter, ok := s.counters[link]
	if ok {
		s.counters[link] = counter + 1
		return nil
	} else {
		return ports.ErrNotExists
	}
}

var _ ports.LinkStorage = &InMemoryLinkStorage{}
