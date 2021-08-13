package services

import (
	"context"
	"fmt"
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	Config         config.DatabaseConfig
	ConnectionPool *pgxpool.Pool
}

func (d *Database) ConnectDatabase() {
	dbpool, err := pgxpool.Connect(context.Background(), d.Config.GetConnectionString())

	d.ConnectionPool = dbpool

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to Database")
}
