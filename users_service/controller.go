package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charleslukes/golang_microservices/helper"
	"go.mongodb.org/mongo-driver/bson"
)

type payload struct {
	Message string `json:"message"`
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

	// add user to the db

	helper.RespondWithJson(w, 200, params)
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
		Message: "User created successfully",
	}

	helper.RespondWithJson(w, 200, payload)
}
