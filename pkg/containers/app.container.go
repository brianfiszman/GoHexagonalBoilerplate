package containers

type AppContainer struct {
	DatabaseContainer *DatabaseContainer
	TicketContainer   TicketContainer
	ServerContainer   ServerContainer
}

func CreateAppContainer() AppContainer {
	var databaseContainer *DatabaseContainer = CreateDatabaseContainer()
	var ticketContainer TicketContainer = CreateTicketContainer(databaseContainer.Database)
	var serverContainer ServerContainer = CreateServerContainer(ticketContainer.Router)

	return AppContainer{
		DatabaseContainer: databaseContainer,
		TicketContainer:   ticketContainer,
		ServerContainer:   serverContainer,
	}
}
