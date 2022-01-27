package adapters

import (
	"context"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type PostgreSQLAdapterMock struct {
	Config         config.DatabaseConfig
	ConnectionPool *pgxpool.Pool
}

func NewPostgreSQLAdapterMock() (adapter *PostgreSQLAdapterMock) {
	adapter = &PostgreSQLAdapterMock{
		Config: config.GetDatabaseConfig(),
	}

	return
}

func (d *PostgreSQLAdapterMock) ConnectDatabase() {
	d.ConnectionPool = &pgxpool.Pool{}

	logrus.Info("Connected to ", d.ConnectionPool.Config().ConnString())
}

func (d *PostgreSQLAdapterMock) Ping() error {
	return d.ConnectionPool.Ping(context.Background())
}
