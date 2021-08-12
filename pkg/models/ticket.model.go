package models

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Ticket struct {
	Id string `json:"id"`
	TicketId string `json:"ticketId"`
	Caller string `json:"caller"`
	Description string `json:"description"`
	ShortDescription string `json:"shortDescription"`
}

func (t *Ticket) Insert (conn *pgxpool.Conn) {
		// Executing SQL query for insertion 
		if _, err := conn.Exec(context.Background(), "INSERT INTO TEST_TABLE(NAME) VALUES($1)", t.Caller); err != nil {
			// Handling error, if occur
			fmt.Println("Unable to insert due to: ", err)
			return
	}

	fmt.Println("Insertion Succesfull")
}
