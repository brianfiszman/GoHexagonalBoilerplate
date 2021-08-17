package containers

import (
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/server"
)

type ServerContainer struct {
	*server.Server
}

func CreateServerContainer(serviceNowRouter routers.ServiceNowRouter) ServerContainer {
	var httpRouter *routers.HTTP_Router = &routers.HTTP_Router{
		ServiceNowRouter: serviceNowRouter,
	}

	httpRouter.ConnectorRouter = httpRouter.NewConnectorRouter()

	var serverContainer ServerContainer = ServerContainer{&server.Server{
		HTTP_Router: httpRouter,
		HTTP_Port:   os.Getenv("HTTP_PORT"),
	}}

	return serverContainer
}
