package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HTTPRouter struct {
	Handler      *chi.Mux
	HealthRouter HealthRouter
}

func (r *HTTPRouter) NewHTTPRouter() *chi.Mux {
	r.Handler = chi.NewRouter()

	r.Handler.Use(middleware.Logger)

	r.Handler.Mount("/auth", NewAuthRouter())
	r.Handler.Mount("/health", r.HealthRouter.Handler)

	return r.Handler
}
