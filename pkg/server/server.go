package server

import (
	"net/http"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/containers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
)

type Server struct {
	HTTP_Router *routers.HTTP_Router
	Database    *services.Database
	HTTP_Port   string
}

func (s *Server) Run() {
	// Connect to Database
	s.Database.ConnectDatabase()

	// Create model
	var ticketModel containers.TicketContainer = containers.CreateTicketContainer(s.Database)

	s.HTTP_Router.ServiceNowRouter = ticketModel.Router

	// Create ConnectorRouter inside Server
	s.HTTP_Router.NewConnectorRouter()

	// Listen on the created router
	http.ListenAndServe(":"+s.HTTP_Port, s.HTTP_Router.ConnectorRouter)
}
