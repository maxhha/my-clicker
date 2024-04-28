package storages

import "github.com/maxhha/my-clicker/internal/interactors/ports"

type RedisLinkStorage struct {
}

func NewRedisLinkStorage() RedisLinkStorage {
	return RedisLinkStorage{}
}

// Create implements ports.LinkStorage.
func (s *RedisLinkStorage) Create(redirect string) (string, error) {
	panic("unimplemented")
}

// GetCounter implements ports.LinkStorage.
func (s *RedisLinkStorage) GetCounter(link string) (uint64, error) {
	panic("unimplemented")
}

// GetRedirect implements ports.LinkStorage.
func (s *RedisLinkStorage) GetRedirect(link string) (string, error) {
	panic("unimplemented")
}

// IncrementCounter implements ports.LinkStorage.
func (s *RedisLinkStorage) IncrementCounter(link string) error {
	panic("unimplemented")
}

var _ ports.LinkStorage = &RedisLinkStorage{}
