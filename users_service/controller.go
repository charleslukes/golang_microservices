package main

import (
	"encoding/json"
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

func signUp(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := SignUp{}

	err := decoder.Decode(&params)

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("error: %v", err))
		return
	}

	// validate
	v_err := params.Validate()

	if v_err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("%v", v_err))
		return
	}

	// find user to the db
	user := &User{}
	email_err := mh.GetOne(user, bson.M{"email": params.Email})

	if email_err == nil {
		helper.RespondWithError(w, 400, "email already exits")
		return
	}

	_, db_err := mh.AddOne(params)

	if db_err != nil {
		fmt.Println(db_err)
		helper.RespondWithError(w, 400, "could not create user")
		return
	}

	payload := payload{
		Message: "signed up successfully",
		Data:    struct{}{},
	}

	helper.RespondWithJson(w, 201, payload)
}

func signIn(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := SignIn{}

	err := decoder.Decode(&params)

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("error: %v", err))
		return
	}

	// validate
	v_err := params.Validate()

	if v_err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("%v", v_err))
		return
	}

	// find user to the db
	user := &User{}
	email_err := mh.GetOne(user, bson.M{"email": params.Email})

	if email_err == nil {
		payload := payload{
			Message: "welcome back!",
			Data:    struct{}{},
		}

		helper.RespondWithJson(w, 200, payload)
	} else {
		helper.RespondWithError(w, 400, "user does not exits")
	}
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
