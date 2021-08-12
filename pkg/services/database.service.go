package services

import (
	"context"
	"fmt"
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectDatabase () {
	dbpool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
		
	}

	defer dbpool.Close()
	
	ticket := models.Ticket{
		Id: "pija",
	}

	fmt.Print(ticket);

	fmt.Println("Connected to Database")

}