package interactors

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/maxhha/my-clicker/internal/interactors/ports"
)

type LinkInteractor struct {
	storage ports.LinkStorage
}

func NewLinkInteractor(storage ports.LinkStorage) LinkInteractor {
	return LinkInteractor{
		storage,
	}
}

var (
	ErrEmptyLink     = errors.New("given link is empty")
	ErrFailParseLink = errors.New("fail to parse link")
)

func (i *LinkInteractor) Create(link string) (string, error) {
	if len(link) == 0 {
		return "", ErrEmptyLink
	}

	_, err := url.ParseRequestURI(link)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrFailParseLink, err)
	}

	var result string
	result, err = i.storage.Create(link)
	if err != nil {
		return "", fmt.Errorf("link storage: %w", err)
	}

	return result, nil
}
