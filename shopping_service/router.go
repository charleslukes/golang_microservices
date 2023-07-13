package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func shoppingRouter() http.Handler  {
	r := chi.NewRouter()

	r.Route("/shopping", func(r chi.Router) {
		// r.Get("/order", getProfile)
	})

	return r
}