package main

import (
	"net/http"
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
)

func main() {
	// Listen
	http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), routers.ConnectorRouter())
}
