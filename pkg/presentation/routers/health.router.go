package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/controllers"
	"github.com/go-chi/chi"
)

type HealthRouter struct {
	Handler          *chi.Mux
	HealthController controllers.HealthController
}

func NewHealthRouter(healthController controllers.HealthController) (h HealthRouter) {
	h = HealthRouter{
		Handler:          chi.NewRouter(),
		HealthController: healthController,
	}

	h.Handler.Get("/", h.HealthController.Health)

	return
}
