package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func shoppingRouter() http.Handler  {
	r := chi.NewRouter()

	r.Route("/shopping", func(r chi.Router) {
		r.Post("/order", createOrder)
		r.Get("/orders", getOrders)
	})

	return r
}