package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type employee struct {
	Name string `validate:"required"`
}

// Error: Key: 'employee.Name' Error:Field validation for 'Name' failed on the 'required' tag
// Notice here that we need to associate meta tags with fields of the struct to let the validator know that you want to validate this field. In the above example, we added the tag with the Name field. This tag is interpreted by the playground validate library.

func main() {
	e := employee{}
	err := validateStruct(e)
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
