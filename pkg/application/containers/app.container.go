package containers

import "github.com/sirupsen/logrus"

type AppContainer struct {
	ServerContainer   ServerContainer
	DatabaseContainer DatabaseContainer
	HealthContainer   HealthContainer
}

func NewAppContainer() (appContainer AppContainer) {
	appContainer.DatabaseContainer = *NewDatabaseContainer()
	logrus.Info("DatabaseContainer Initialized")

	appContainer.HealthContainer = NewHealthContainer(appContainer.DatabaseContainer.Database)
	logrus.Info("HealthContainer Initialized")

	appContainer.ServerContainer = NewServerContainer(appContainer.HealthContainer.Router)
	logrus.Info("ServerContainer Initialized")

	return
}
