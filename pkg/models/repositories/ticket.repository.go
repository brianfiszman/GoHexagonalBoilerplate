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

func (t *TicketRepository) Insert(ticket schemas.Ticket) {
	// Executing SQL query for insertion
	if _, err := t.Database.ConnectionPool.Exec(context.Background(), "INSERT INTO TEST_TABLE(NAME) VALUES($1)", ticket.Caller); err != nil {
		// Handling error, if occur
		fmt.Println("Unable to insert due to: ", err)
		return
	}

	fmt.Println("Insertion Succesfull")
}
