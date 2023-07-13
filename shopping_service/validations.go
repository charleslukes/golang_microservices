package main

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type CreateOrder struct {
	Shop
}

func (v CreateOrder) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.OrderId, validation.Required),
		validation.Field(&v.CustomerId, validation.Required),
		validation.Field(&v.Amount, validation.Required),
		validation.Field(&v.Status, validation.Required),
	)
}
