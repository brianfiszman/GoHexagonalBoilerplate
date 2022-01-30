package containers

import (
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/domain/interfaces"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/infrastructure/adapters"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/infrastructure/config"
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
