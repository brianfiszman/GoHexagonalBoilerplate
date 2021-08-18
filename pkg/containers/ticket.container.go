package containers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/repositories"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/services"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	infrastructure "github.com/brianfiszman/GoFromZeroToHero/pkg/services"
)

// Container
type TicketContainer struct {
	Repository repositories.TicketRepository
	Service    services.TicketService
	Controller controllers.TicketController
	Router     routers.ServiceNowRouter
}

func CreateTicketContainer(d *infrastructure.Database) TicketContainer {
	var ticketRepository = repositories.TicketRepository{Database: d}
	var ticketService = services.TicketService{Repository: ticketRepository, ServiceNowConfig: config.LoadServiceNowConfig()}
	var t TicketContainer = TicketContainer{
		Repository: ticketRepository,
		Service:    ticketService,
		Controller: controllers.TicketController{
			Service: ticketService,
		},
	}

	t.Router.Controller = t.Controller
	t.Router.Router = t.Router.NewServiceNowRouter()

	return t
}
