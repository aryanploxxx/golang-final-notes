package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type employee struct {
	Age int `validate:"required,gte=10,lte=20"`
}

// Validate range of interger in Go

func main() {
	e := employee{}
	err := validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	e = employee{Age: 5}
	err = validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	e = employee{Age: 25}
	err = validateStruct(e)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func validateStruct(e employee) error {
	validate = validator.New()
	err := validate.Struct(e)
	if err != nil {
		return err
	}
	return nil
}
