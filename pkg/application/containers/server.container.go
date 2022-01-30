package containers

import (
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/infrastructure/http"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/presentation/routers"
)

type ServerContainer struct {
	Server *http.Server
}

func NewServerContainer(healthRouter routers.HealthRouter) (serverContainer ServerContainer) {
	var httpRouter *routers.HTTPRouter = &routers.HTTPRouter{HealthRouter: healthRouter}

	httpRouter.Handler = httpRouter.NewHTTPRouter()

	serverContainer = ServerContainer{
		Server: http.NewServer(httpRouter),
	}

	return
}
