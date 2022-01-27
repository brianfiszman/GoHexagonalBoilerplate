package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/controllers"
	"github.com/go-chi/chi"
)

type HealthRouter struct {
	Handler          *chi.Mux
	HealthController controllers.HealthController
}

func (r HealthRouter) NewHealthRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", r.HealthController.IsHealthy)

	return router
}
