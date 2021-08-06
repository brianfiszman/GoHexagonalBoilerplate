package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get(
		"/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("Welcome"))
		})

	http.ListenAndServe(":3000", router)
}
