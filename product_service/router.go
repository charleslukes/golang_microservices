package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
)

func productRouter() http.Handler  {
	r := chi.NewRouter()

	r.Route("/product", func(r chi.Router) {
		r.Get("/", getProfile)
	})

	return r
}