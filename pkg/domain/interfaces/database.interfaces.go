package interfaces

import "github.com/jackc/pgx/v4/pgxpool"

type Database interface {
	ConnectDatabase()
	GetConnection() *pgxpool.Pool
}
