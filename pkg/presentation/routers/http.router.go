package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HTTPRouter struct {
	Handler *chi.Mux
}

func (r *HTTPRouter) NewHTTPRouter() *chi.Mux {
	r.Handler = chi.NewRouter()

	r.Handler.Use(middleware.Logger)

	r.Handler.Mount("/auth", NewAuthRouter())
	r.Handler.Get("/", controllers.GetHeartBeat)

	return r.Handler
}
