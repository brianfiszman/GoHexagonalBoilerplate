package main

import (
	"net/http"
	"os"

	"github.com/brianfiszman/GoFromZeroToHero/pkg/config"
	"github.com/brianfiszman/GoFromZeroToHero/pkg/routers"
)

func main() {
	// Load Environment Variables
	config.LoadEnvironmentVariables()

	// Listen
	http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), routers.Router)
}
