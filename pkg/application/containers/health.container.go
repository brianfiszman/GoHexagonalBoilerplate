package containers

import (
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/application/controllers"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/domain/interfaces"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/domain/services"
	"github.com/brianfiszman/GoHexagonalBoilerplate/pkg/presentation/routers"
)

type HealthContainer struct {
	Router     routers.HealthRouter
	Controller controllers.HealthController
	Service    services.HealthService
}

func NewHealthContainer(d interfaces.Database) (h HealthContainer) {
	var healthService services.HealthService = services.HealthService{
		Database: d,
	}

	var healthController controllers.HealthController = controllers.HealthController{
		HealthService: healthService,
	}

	var healthRouter routers.HealthRouter = routers.NewHealthRouter(healthController)

	h.Router = healthRouter
	h.Controller = healthController
	h.Service = healthService

	return
}
