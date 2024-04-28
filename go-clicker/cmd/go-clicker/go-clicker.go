package main

import (
	"github.com/maxhha/my-clicker/internal/configuration"
	"github.com/maxhha/my-clicker/internal/server"
)

func main() {
	configuration := configuration.NewConfigurationBuilder().
		SetAddr(":3000").Build()

	server := server.NewServer(configuration)
	server.Run()
}
