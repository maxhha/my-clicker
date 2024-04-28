package server

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/maxhha/my-clicker/internal/interactors"
	"github.com/maxhha/my-clicker/internal/interactors/ports"
)

func (s *Server) HandleLinkCreate(c fiber.Ctx) error {
	var (
		err  error
		data struct {
			Redirect string
		}
	)
	if err = c.Bind().Body(&data); err != nil {
		return err
	}

	var link string
	if link, err = s.linkInteractor.Create(data.Redirect); err != nil {
		if errors.Is(err, interactors.ErrEmptyLink) || errors.Is(err, interactors.ErrFailParseLink) {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		return err
	}

	var resp = struct {
		LinkId string `json:"link_id"`
		Url    string `json:"url"`
	}{
		LinkId: link,
		Url:    fmt.Sprintf("%s/l/%s", s.configuration.ServerUrl(), link),
	}

	return c.Status(200).JSON(resp)
}

func (s *Server) HandleLinkGetCounter(c fiber.Ctx) error {
	link := c.Params("link")
	counter, err := s.linkInteractor.GetCounter(link)
	if err != nil {
		if errors.Is(err, ports.ErrNotExists) || errors.Is(err, interactors.ErrEmptyLink) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return err
	}

	var resp = struct {
		LinkId string `json:"link_id"`
		Total  uint64 `json:"total"`
	}{
		LinkId: link,
		Total:  counter,
	}

	return c.Status(200).JSON(resp)
}

func (s *Server) HandleLinkRedirect(c fiber.Ctx) error {
	link := c.Params("link")
	redirect_url, err := s.linkInteractor.Visit(link)
	if err != nil {
		if errors.Is(err, ports.ErrNotExists) || errors.Is(err, interactors.ErrEmptyLink) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return err
	}

	return c.Redirect().To(redirect_url)
}
