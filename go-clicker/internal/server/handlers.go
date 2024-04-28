package server

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/maxhha/my-clicker/internal/interactors"
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
