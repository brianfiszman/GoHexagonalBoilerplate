package containers

import (
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/server"
	"github.com/go-chi/chi"
)

type AppContainer struct {
	DatabaseContainer *DatabaseContainer
	TicketContainer   TicketContainer
	Server            *server.Server
}

func CreateAppContainer() AppContainer {
	var databaseContainer *DatabaseContainer = CreateDatabaseContainer()
	var ticketContainer TicketContainer = CreateTicketContainer(databaseContainer.Database)
	var server *server.Server = &server.Server{
		HTTP_Router: &routers.HTTP_Router{
			ConnectorRouter:  chi.NewRouter(),
			ServiceNowRouter: ticketContainer.Router,
		},
		HTTP_Port: os.Getenv("HTTP_PORT"),
	}

	return AppContainer{
		DatabaseContainer: databaseContainer,
		TicketContainer:   ticketContainer,
		Server:            server,
	}
}
