package containers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/controllers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/domain/interfaces"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/domain/services"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/presentation/routers"
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
