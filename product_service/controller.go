package main

import (
	"encoding/json"
	"fmt"
	"github.com/charleslukes/golang_microservices/helper"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

type payload struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := &Product{}

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

	// check if product name already exits
	p_err := mh.GetOne(params, bson.M{"name": params.Name})

	if p_err == nil {
		params.Name = strings.ToLower(params.Name)
		res, err := mh.AddOne(params)

		if err != nil {
			helper.RespondWithError(w, 400, fmt.Sprintf("%v", v_err))
			return
		}

		payload := payload{
			Message: "product created",
			Data:    res,
		}
		helper.RespondWithJson(w, 201, payload)
	} else {
		helper.RespondWithError(w, 400, fmt.Sprintf("Product with name:`%v` already exits", params.Name))
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := mh.Get(struct{}{})
	helper.RespondWithJson(w, 201, products)
}
