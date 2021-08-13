package routers

import (
	"github.com/go-chi/chi"
)

type HTTP_Router struct {
	ConnectorRouter *chi.Mux
	ServiceNowRouter ServiceNowRouter
}

func (r *HTTP_Router) NewConnectorRouter() {
	r.ConnectorRouter.Mount("/auth", NewAuthRouter())
	r.ConnectorRouter.Mount("/tickets", r.ServiceNowRouter.NewServiceNowRouter())
}
