package storages

import (
	"errors"
	"fmt"
	"sync"

	"github.com/maxhha/my-clicker/internal/interactors/ports"
	"github.com/teris-io/shortid"
)

// Only for simple tests
type InMemoryLinkStorage struct {
	createTries   int
	capacityLimit int
	linksMx       sync.RWMutex
	links         map[string]string
	countersMx    sync.Mutex
	counters      map[string]uint64
}

func NewInMemoryLinkStorage() InMemoryLinkStorage {
	return InMemoryLinkStorage{
		createTries:   10,
		capacityLimit: 10,
		linksMx:       sync.RWMutex{},
		links:         make(map[string]string),
		countersMx:    sync.Mutex{},
		counters:      make(map[string]uint64),
	}
}

var ErrFailGenerate = errors.New("fail generate unique short id")
var ErrFull = errors.New("storage is full")

// Create implements ports.LinkStorage.
func (s *InMemoryLinkStorage) Create(redirect string) (string, error) {
	var (
		id  string
		err error
	)

	s.linksMx.Lock()
	defer s.linksMx.Unlock()

	if len(s.links) > s.capacityLimit {
		return "", ErrFull
	}

	for try := 0; len(id) == 0 && try < s.createTries; try++ {
		id, err = shortid.Generate()
		if err != nil {
			return "", fmt.Errorf("shortid.Generate: %w", err)
		}

		_, exists := s.links[id]
		if exists {
			id = ""
		}
	}

	if len(id) == 0 {
		return "", ErrFailGenerate
	}

	s.links[id] = redirect
	s.countersMx.Lock()
	defer s.countersMx.Unlock()
	s.counters[id] = 0

	return id, nil
}

// GetCounter implements ports.LinkStorage.
func (s *InMemoryLinkStorage) GetCounter(link string) (uint64, error) {
	s.countersMx.Lock()
	defer s.countersMx.Unlock()
	counter, ok := s.counters[link]
	if ok {
		return counter, nil
	} else {
		return 0, ports.ErrNotExists
	}
}

// GetRedirect implements ports.LinkStorage.
func (s *InMemoryLinkStorage) GetRedirect(link string) (string, error) {
	s.linksMx.RLock()
	defer s.linksMx.RUnlock()
	redirect, ok := s.links[link]
	if ok {
		return redirect, nil
	} else {
		return "", ports.ErrNotExists
	}
}

// IncrementCounter implements ports.LinkStorage.
func (s *InMemoryLinkStorage) IncrementCounter(link string) error {
	s.countersMx.Lock()
	defer s.countersMx.Unlock()
	counter, ok := s.counters[link]
	if ok {
		s.counters[link] = counter + 1
		return nil
	} else {
		return ports.ErrNotExists
	}
}

var _ ports.LinkStorage = &InMemoryLinkStorage{}
