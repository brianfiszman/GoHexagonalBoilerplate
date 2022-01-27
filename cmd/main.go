package main

import "github.com/brianfiszman/GoFromZeroToHero/pkg/containers"

func main() {
	var app containers.AppContainer = containers.NewAppContainer()

	// Server initializes listening
	app.ServerContainer.Server.Run()
}
