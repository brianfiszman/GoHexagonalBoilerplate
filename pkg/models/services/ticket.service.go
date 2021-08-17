package services

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/repositories"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/schemas"
)

type TicketService struct {
	Repository repositories.TicketRepository
}

// Creates a ticket in DB
func (s TicketService) Create(body map[string]string) {
	var ticket schemas.Ticket = schemas.Ticket{Caller: body["caller"]}

	s.Repository.Insert(ticket)
}
