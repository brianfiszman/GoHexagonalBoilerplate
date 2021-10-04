package containers

import (
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/server"
)

type ServerContainer struct {
	*server.Server
}

func CreateServerContainer() (serverContainer ServerContainer) {
	var httpRouter *routers.HTTP_Router = &routers.HTTP_Router{}

	httpRouter.ConnectorRouter = httpRouter.NewConnectorRouter()

	serverContainer = ServerContainer{
		&server.Server{
			HTTP_Router: httpRouter,
			HTTP_Port:   os.Getenv("HTTP_PORT"),
		},
	}

	return
}
