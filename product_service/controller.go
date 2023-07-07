package main

import (
	"fmt"
	"net/http"

	"github.com/charleslukes/golang_microservices/helper"
	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"go.mongodb.org/mongo-driver/bson"
)

type payload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	emailAddress := chi.URLParam(r, "emailAddress")
	err := validation.Validate(emailAddress, is.Email)

	if err != nil {
		helper.RespondWithError(w, 404, fmt.Sprintf("error: %v", err))
		return
	}

	user := &User{}
	err = mh.GetOne(user, bson.M{"email": emailAddress})
	if err != nil {
		helper.RespondWithError(w, 404, fmt.Sprintf("error: %v", err))
		return
	}
	payload := payload{
		Message: "successful",
		Data:    user,
	}

	helper.RespondWithJson(w, 200, payload)
}
