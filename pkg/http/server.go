package http

import (
	"net/http"
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Router    *routers.HTTPRouter
	HTTP_Port string
}

func NewServer(httpRouter *routers.HTTPRouter) (server *Server) {
	return &Server{
		Router:    httpRouter,
		HTTP_Port: os.Getenv("HTTP_PORT"),
	}
}

func (s *Server) Run() {
	// Listen on the created router
	logrus.Info("Listening on port: ", s.HTTP_Port)

	err := http.ListenAndServe(":"+s.HTTP_Port, s.Router.Handler)

	logrus.Fatal(err)
}
