package containers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
)

type DatabaseContainer struct {
	DatabaseConfig config.DatabaseConfig
	Database       *services.Database
}

func CreateDatabaseContainer() *DatabaseContainer {
	var d *DatabaseContainer = &DatabaseContainer{
		DatabaseConfig: config.GetDatabaseConfig(),
		Database: &services.Database{
			Config: config.GetDatabaseConfig(),
		},
	}

	d.Database.ConnectDatabase()

	return d
}
