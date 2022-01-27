package containers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/http"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
)

type ServerContainer struct {
	Server *http.Server
}

func NewServerContainer() (serverContainer ServerContainer) {
	var httpRouter *routers.HTTPRouter = &routers.HTTPRouter{}

	httpRouter.Handler = httpRouter.NewHTTPRouter()

	serverContainer = ServerContainer{
		Server: http.NewServer(httpRouter),
	}

	return
}
