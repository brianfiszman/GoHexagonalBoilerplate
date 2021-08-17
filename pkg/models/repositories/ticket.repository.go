package repositories

import (
	"context"
	"fmt"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/models/schemas"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
)

type TicketRepository struct {
	Database *services.Database
}

func (t *TicketRepository) Insert(ticket schemas.ServiceNowResultDTO) {
	// Executing SQL query for insertion
	if _, err := t.Database.ConnectionPool.Exec(context.Background(), 
	"INSERT INTO model_dev.ticket(assignee, ticketing_system, created_at, issue_type, project_key, project_name, summary, ticketing_system_ticket_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8)", 
	ticket.Result.Assignee, "SERVICENOW", ticket.Result.CreatedAt, ticket.Result.IssueType, ticket.Result.ProjectKey, ticket.Result.ProjectName, ticket.Result.ShortDescription, ticket.Result.TicketId); err != nil {
		// Handling error, if occur
		fmt.Println("Unable to insert due to: ", err)
		return
	}

	fmt.Println("Insertion Succesfull")
}
