package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/controllers"
	"github.com/go-chi/chi"
)

type HealthRouter struct {
	HealthController controllers.HealthController
}

func (r HealthRouter) NewHealthRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/", r.HealthController.IsHealthy)

	return router
}
