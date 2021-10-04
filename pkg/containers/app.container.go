package containers

type AppContainer struct {
	ServerContainer ServerContainer
}

func CreateAppContainer() (appContainer AppContainer) {
	appContainer.ServerContainer = CreateServerContainer()

	return
}
