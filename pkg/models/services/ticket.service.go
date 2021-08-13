package services

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/repositories"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/schemas"
)

type TicketService struct {
	Repository repositories.TicketRepository
}

// Creates a ticket in DB
func (s TicketService) Create(ticket schemas.Ticket) {
	s.Repository.Insert(ticket)
}
