package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
)

type ServiceNowRouter struct {
	router	*chi.Mux
	controller	controllers.TicketController
}

func (r ServiceNowRouter) NewServiceNowRouter() *chi.Mux {
	r.router = chi.NewRouter()

	r.router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(services.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Get("/users", controllers.GetUsersList)
		router.Get("/", controllers.GetTicketList)
		router.Post("/", r.controller.CreateTicket)
	})

	return r.router
}
