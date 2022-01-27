package containers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/controllers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/domain/interfaces"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/domain/services"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/presentation/routers"
)

type HealthContainer struct {
	Database   interfaces.Database
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

	var healthRouter routers.HealthRouter = routers.HealthRouter{
		HealthController: healthController,
	}

	healthRouter.Handler = healthRouter.NewHealthRouter()

	h.Router = healthRouter
	h.Controller = healthController
	h.Service = healthService

	return
}
