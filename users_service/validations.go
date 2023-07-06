package main

import (
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type SignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (v SignUp) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&v.Email, validation.Required, is.Email),
		validation.Field(&v.Password, validation.Required),
	)
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (v SignIn) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required, is.Email),
		validation.Field(&v.Password, validation.Required),
	)
}

type Address struct {
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	Country    string `json:"country"`
}

func (v Address) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.PostalCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		validation.Field(&v.City, validation.Required, validation.Length(5, 50)),
		validation.Field(&v.Country, validation.Required, validation.Length(5, 50)),
	)
}
