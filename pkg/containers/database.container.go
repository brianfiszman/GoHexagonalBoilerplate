package containers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/adapters"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/interfaces"
)

type DatabaseContainer struct {
	DatabaseConfig config.DatabaseConfig
	Database       interfaces.Database
}

func NewDatabaseContainer() *DatabaseContainer {
	var d *DatabaseContainer = &DatabaseContainer{
		DatabaseConfig: config.GetDatabaseConfig(),
		Database:       adapters.NewPostgreSQLAdapter(),
	}

	d.Database.ConnectDatabase()

	return d
}
