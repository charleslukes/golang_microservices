package main

import (
	"github.com/charleslukes/golang_microservices/helper"
	"net/http"
)

func handleError(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithError(w, 400, "something went wrong")
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJson(w, 200, struct{}{})
}
