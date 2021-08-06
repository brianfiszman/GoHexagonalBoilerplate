package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
)

var (
	Router chi.Router = CreateRouter()
)

func CreateRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/tickets", controllers.GetTicketList)
	router.Post("/tickets", controllers.CreateTicket)
	return router
}
