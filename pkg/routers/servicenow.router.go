package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
)

func NewServiceNowRouter() chi.Router {
	router := chi.NewRouter()

	router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(services.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Get("/users", controllers.GetUsersList)
		router.Get("/", controllers.GetTicketList)
		router.Post("/", controllers.CreateTicket)
	})

	return router
}
