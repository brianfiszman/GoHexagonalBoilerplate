package main

import (
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/server"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
	"github.com/go-chi/chi"
)

func main() {
	var database *services.Database = &services.Database{
		Config: config.GetDatabaseConfig(),
	}

	var httpRouter *routers.HTTP_Router = &routers.HTTP_Router{
		ConnectorRouter: chi.NewRouter(),
	}

	var srv *server.Server = &server.Server{
		HTTP_Router: httpRouter,
		Database:    database,
		HTTP_Port:   os.Getenv("HTTP_PORT"),
	}

	// Server initializes listening
	srv.Run()
}
