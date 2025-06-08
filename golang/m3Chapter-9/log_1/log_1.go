package main

import (
	"errors"
	"log"
	"reflect"
)

// Alpha struct for string operations
type Alpha struct {
}

// Numeric struct for integer operations
type Numeric struct {
}

// Add method for Alpha struct concatenates two strings
// Returns the concatenated string and any error that occurred
func (a Alpha) Add(x string, y string) (string, error) {
	var err error
	// Check the types of input parameters using reflection
	xType := reflect.TypeOf(x).Kind()
	yType := reflect.TypeOf(y).Kind()
	// Validate that both inputs are strings
	if xType != reflect.String || yType != reflect.String {
		err = errors.New("incorrect type for strings a or b")
	}
	finalString := x + y
	return finalString, err
}

// Add method for Numeric struct adds two integers
// Returns the sum and any error that occurred
func (n Numeric) Add(x int, y int) (int, error) {
	var err error

	// Check the types of input parameters using reflection
	xType := reflect.TypeOf(x).Kind()
	yType := reflect.TypeOf(y).Kind()
	// Validate that both inputs are integers
	if xType != reflect.Int || yType != reflect.Int {
		err = errors.New("incorrect type for integer a or b")
	}
	return x + y, err
}

func main() {
	// Create instances of Numeric and Alpha structs
	n1 := Numeric{}
	a1 := Alpha{}

	// Test numeric addition
	z, err := n1.Add(5, 2)
	if err != nil {
		log.Println("Error", err)
	}
	log.Println(z)

	// Test string concatenation
	y, err := a1.Add("super", "lative")
	if err != nil {
		log.Println("Error", err)
	}
	log.Println(y)
}
