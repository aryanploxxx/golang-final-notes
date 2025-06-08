package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "true"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	}

	input = "false"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	}

	input = "garbage"
	if val, err := strconv.ParseBool(input); err == nil {
		fmt.Printf("%T, %v\n", val, val)
	} else {
		fmt.Println("Given input is not a bool")
	}
}

// bool, true
// bool, false
// Given input is not a bool
// -> strconv.ParseBool() function can be used to parse a string representation of a bool.

// Different format specifiers can be used to print a boolean in either bool or string.

// %t can be used to print in the boolean form
// %v will print the default string. “true” for true and “false” for false
