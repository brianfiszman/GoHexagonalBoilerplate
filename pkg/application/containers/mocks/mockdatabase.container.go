package mocks

import (
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/domain/interfaces"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/infrastructure/adapters/mocks"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/infrastructure/config"
)

type MockDatabaseContainer struct {
	DatabaseConfig config.DatabaseConfig
	Database       interfaces.Database
}

func NewMockDatabaseContainer() *MockDatabaseContainer {
	var d *MockDatabaseContainer = &MockDatabaseContainer{
		DatabaseConfig: config.GetDatabaseConfig(),
		Database:       mocks.NewMockPostgreSQLAdapter(),
	}

	d.Database.ConnectDatabase()

	return d
}
