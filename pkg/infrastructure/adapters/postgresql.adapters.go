package adapters

import (
	"context"

	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type PostgreSQLAdapter struct {
	Config         config.DatabaseConfig
	ConnectionPool *pgxpool.Pool
}

func NewPostgreSQLAdapter() (adapter *PostgreSQLAdapter) {
	adapter = &PostgreSQLAdapter{
		Config: config.GetDatabaseConfig(),
	}

	return
}

func (d *PostgreSQLAdapter) ConnectDatabase() {
	dbpool, err := pgxpool.Connect(context.Background(), d.Config.GetConnectionString())

	d.ConnectionPool = dbpool

	if err != nil {
		logrus.Fatal("Unable to connect to database: ", err)
	}

	logrus.Info("Connected to ", d.ConnectionPool.Config().ConnString())
}

func (d *PostgreSQLAdapter) Ping() error {
	return d.ConnectionPool.Ping(context.Background())
}
