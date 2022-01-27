package services

import (
	"encoding/json"
	"fmt"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/dtos"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/domain/schemas"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/infrastructure/config"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/infrastructure/repositories"
	"github.com/go-resty/resty/v2"
)

type TicketService struct {
	Repository repositories.TicketRepository
	config.ServiceNowConfig
}

var restClient resty.Client = *resty.New()

/*
* Send a REST Request to ServiceNow API and
* fetches the user ticket lists on success
 */
func (s TicketService) GetTickets() (*resty.Response, error) {
	res, err := restClient.
		R().
		EnableTrace().
		SetBasicAuth(s.USER, s.PASS).
		Get(s.API_URL + "/now/table/incident")

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, err
}

/*
* Send a REST Request to ServiceNow API and
* fetches the users list on success
 */
func (s TicketService) GetUsers() (*resty.Response, error) {
	res, err := restClient.
		R().
		EnableTrace().
		SetBasicAuth(s.USER, s.PASS).
		Get(s.API_URL + "/now/table/sys_user")

	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, err
}

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
