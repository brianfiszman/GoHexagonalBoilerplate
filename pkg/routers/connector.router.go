package routers

import (
	"github.com/go-chi/chi"
)

func ConnectorRouter() chi.Router {
	router := chi.NewRouter()
	router.Mount("/auth", CreateAuthRouter())
	router.Mount("/tickets", CreateServiceNowRouter())
	return router
}