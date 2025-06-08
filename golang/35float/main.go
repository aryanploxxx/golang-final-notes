package main

import (
	"fmt"
	"strconv"
)

func main() {
	e1 := "1.3434"
	if s, err := strconv.ParseFloat(e1, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat(e1, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}

// strconv.ParseFloat() function can be used to parse a string representation of a float
// func ParseFloat(s string, bitSize int) (float64, error)

// Some points worth noting:
// The first argument is the string representation of a float
// The second argument is the bitSize which specifies the precision. It is 32 for float32 and 64 for float64
// The return value will always be float64 but will be convertible to float32 without any change in its value.

// float64, 1.343400001525879
// float64, 1.3434
