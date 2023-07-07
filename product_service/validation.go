package main

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (v Product) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&v.Desc, validation.Required, validation.Length(3, 100)),
		validation.Field(&v.Type, validation.Required, validation.Length(3, 50)),
		validation.Field(&v.Unit, validation.Required, is.UTFNumeric),
		validation.Field(&v.Price, validation.Required, is.UTFNumeric),
		validation.Field(&v.Available, validation.Required),
		validation.Field(&v.Supplier, validation.Required),
	)
}
