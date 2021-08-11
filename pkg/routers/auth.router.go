package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
	"github.com/go-chi/chi"
)

func CreateAuthRouter() chi.Router{
	router := chi.NewRouter()

	router.Post("/", controllers.Auth)
	
	return router
}