package interactors

import (
	"errors"
	"fmt"
	"net/url"

	"log"

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

func (i *LinkInteractor) Visit(link string) (string, error) {
	if len(link) == 0 {
		return "", ErrEmptyLink
	}

	redirect, err := i.storage.GetRedirect(link)
	if err != nil {
		return "", fmt.Errorf("link storage: %w", err)
	}

	// TODO: run in goroutine?
	err = i.storage.IncrementCounter(link)
	if err != nil {
		log.Printf("LinkInteractor.Visit(%s): increment counter: %v\n", link, err)
	}

	return redirect, nil
}

func (i *LinkInteractor) GetCounter(link string) (uint64, error) {
	if len(link) == 0 {
		return 0, ErrEmptyLink
	}

	counter, err := i.storage.GetCounter(link)
	if err != nil {
		return 0, fmt.Errorf("link storage: %w", err)
	}

	return counter, nil
}
