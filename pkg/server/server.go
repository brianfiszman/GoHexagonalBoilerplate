package server

import (
	"net/http"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
)

type Server struct {
	HTTP_Router *routers.HTTP_Router
	HTTP_Port   string
}

func (s *Server) Run() {
	// Listen on the created router
	http.ListenAndServe(":"+s.HTTP_Port, s.HTTP_Router.ConnectorRouter)
}
