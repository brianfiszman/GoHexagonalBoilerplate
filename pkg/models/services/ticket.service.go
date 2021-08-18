package services

import (
	"encoding/json"
	"fmt"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/dtos"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/repositories"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/schemas"
	"github.com/go-resty/resty/v2"
)

type TicketService struct {
	Repository repositories.TicketRepository
	config.ServiceNowConfig
}

var restClient resty.Client = *resty.New()

/*
* Send a REST Request to ServiceNow API and generates
* a Ticket in PostgreSQL Database on Success
 */
func (s TicketService) Create(createTicketDTO dtos.CreateTicketDTO) (*resty.Response, error) {
	res, err := restClient.
		R().
		EnableTrace().
		SetBasicAuth(s.USER, s.PASS).
		SetBody(createTicketDTO).
		Post(s.API_URL + "/now/table/incident")

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	ticket, err := createResultDTOToTicketSchema(res)

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	s.Repository.Insert(ticket)

	return res, err
}

// This function serializes the body response into a schemas.Ticket type
func createResultDTOToTicketSchema(res *resty.Response) (schemas.Ticket, error) {
	createResultDto := dtos.CreateServiceNowResultDTO{}
	ticket := schemas.Ticket{}

	// Serialize the response body into the CreateResultDTO
	err := json.Unmarshal([]byte(res.Body()), &createResultDto)

	if err != nil {
		fmt.Println(err)
		return ticket, err
	}

	// Pass the CreateResultDTO to the schemas.Ticket
	ticket = createResultDto.Result

	return ticket, err
}
