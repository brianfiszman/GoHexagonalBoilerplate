package services

import (
	"encoding/json"
	"fmt"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/repositories"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/schemas"
	"github.com/go-resty/resty/v2"
)

type TicketService struct {
	Repository repositories.TicketRepository
}

// Creates a ticket in DB
func (s TicketService) Create(body *resty.Response) {

	var ticket schemas.ServiceNowResultDTO = schemas.ServiceNowResultDTO{}
	err := json.Unmarshal([]byte(body.Body()), &ticket)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.Repository.Insert(ticket)
}
