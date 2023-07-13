package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func productRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/product", func(r chi.Router) {
		r.Get("/{id}", getProductById)
		r.Get("/all-products", getProducts)
		r.Post("/", createProduct)
	})

	return r
}
