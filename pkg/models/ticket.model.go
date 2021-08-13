package models

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/repositories"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/services"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	infrastructure "github.com/brianfiszman/GoFromZeroToHero/pkg/services"
)

// Container
type TicketModel struct {
	Repository	repositories.TicketRepository
	Service		  services.TicketService
	Controller  controllers.TicketController
	router			routers.ServiceNowRouter
}

func CreateTicketContainer(d *infrastructure.Database){
	
	var ticketRepository = repositories.TicketRepository{Database: d}

	var ticketService = services.TicketService{Repository: ticketRepository}

	var t TicketModel = TicketModel{
		Repository: ticketRepository,
		Service: ticketService,
		Controller: controllers.TicketController{
			Service: ticketService,
		},
	}

	t.router = routers.ServiceNowRouter{}

}