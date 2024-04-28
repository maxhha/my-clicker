package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/maxhha/my-clicker/internal/configuration"
	"github.com/maxhha/my-clicker/internal/interactors"
	"github.com/maxhha/my-clicker/internal/storages"
)

type Server struct {
	configuration configuration.Configuartion

	linkInteractor *interactors.LinkInteractor
	app            *fiber.App
}

func NewServer(configuration configuration.Configuartion) Server {
	app := fiber.New(fiber.Config{
		// DisableKeepalive: true,
	})

	linkStorage := storages.NewInMemoryLinkStorage()

	linkInteractor := interactors.NewLinkInteractor(&linkStorage)

	server := Server{
		configuration:  configuration,
		linkInteractor: &linkInteractor,
		app:            app,
	}

	app.Post("/api/link/create", server.HandleLinkCreate)
	app.Get("/api/link/:link/counter", server.HandleLinkGetCounter)
	app.Get("/l/:link", server.HandleLinkRedirect)

	return server
}

func (s *Server) Run() {
	log.Fatal(s.app.Listen(s.configuration.Addr()))
}
