package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func userRouter() http.Handler  {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/{emailAddress}", getProfile)
		r.Post("/signin",  signIn)
		r.Post("/signup",  signUp)
	})

	return r
}