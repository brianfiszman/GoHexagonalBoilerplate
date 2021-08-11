package routers

import (
	"github.com/go-chi/chi"
)

func TicketingRouter() chi.Router {
	router := chi.NewRouter()
	router.Mount("/auth", CreateAuthRouter())
	router.Mount("/tickets", CreateServiceNowRouter())
	return router
}