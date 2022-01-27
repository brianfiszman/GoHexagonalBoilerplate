package containers

import "github.com/sirupsen/logrus"

type AppContainer struct {
	ServerContainer   ServerContainer
	DatabaseContainer DatabaseContainer
}

func NewAppContainer() (appContainer AppContainer) {
	appContainer.DatabaseContainer = *NewDatabaseContainer()
	logrus.Info("DatabaseContainer Initialized")

	appContainer.ServerContainer = NewServerContainer()
	logrus.Info("ServerContainer Initialized")

	return
}
