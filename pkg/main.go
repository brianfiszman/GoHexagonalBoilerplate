package main

import (
	"net/http"
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/services"
)

func main() {
	// Connects DataBase
	services.ConnectDatabase()

	// Listen
	http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), routers.ConnectorRouter())
}
