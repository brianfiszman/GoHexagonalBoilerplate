package mocks

import (
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/application/containers"
	"github.com/sirupsen/logrus"
)

type MockAppContainer struct {
	MockDatabaseContainer MockDatabaseContainer
	ServerContainer       containers.ServerContainer
	HealthContainer       containers.HealthContainer
}

func NewMockAppContainer() (appContainer MockAppContainer) {
	appContainer.MockDatabaseContainer = *NewMockDatabaseContainer()
	logrus.Info("MockDatabaseContainer Initialized")

	appContainer.HealthContainer = containers.NewHealthContainer(appContainer.MockDatabaseContainer.Database)
	logrus.Info("HealthContainer Initialized")

	appContainer.ServerContainer = containers.NewServerContainer(appContainer.HealthContainer.Router)
	logrus.Info("ServerContainer Initialized")

	return
}
