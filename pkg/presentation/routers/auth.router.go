package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/application/controllers"
	"github.com/go-chi/chi"
)

func NewAuthRouter() chi.Router {
	router := chi.NewRouter()

	router.Post("/", controllers.Auth)

	return router
}
