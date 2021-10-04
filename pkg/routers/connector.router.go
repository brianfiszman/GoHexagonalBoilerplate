package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HTTP_Router struct {
	ConnectorRouter *chi.Mux
}

func (r *HTTP_Router) NewConnectorRouter() *chi.Mux {
	r.ConnectorRouter = chi.NewRouter()

	r.ConnectorRouter.Use(middleware.Logger)

	r.ConnectorRouter.Mount("/auth", NewAuthRouter())
	r.ConnectorRouter.Get("/", controllers.GetHeartBeat)

	return r.ConnectorRouter
}
