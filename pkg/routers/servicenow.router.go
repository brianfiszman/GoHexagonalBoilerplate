package routers

import (
	"github.com/brianfiszman/GoFromZeroToHero/pkg/controllers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
)

type ServiceNowRouter struct {
	Router     *chi.Mux
	Controller controllers.TicketController
}

func (r ServiceNowRouter) NewServiceNowRouter() *chi.Mux {
	r.Router = chi.NewRouter()

	r.Router.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(services.TokenAuth))
		router.Use(jwtauth.Authenticator)
		router.Get("/users", controllers.GetUsersList)
		router.Get("/", r.Controller.GetTicketList)
		router.Post("/", r.Controller.CreateTicket)
	})

	return r.Router
}
