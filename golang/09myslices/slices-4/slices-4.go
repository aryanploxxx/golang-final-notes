package main

import "fmt"

func main() {
	sample := "Hello"
	suffix := "World"

	result := append([]byte(sample), suffix...)
	fmt.Printf("sample: %s\n", string(result))

	// result := append(sample, suffix...)
	// this will give error -> ./slices-4.go:33:18: first argument to append must be slice; have string
	// A string in go is nothing but a sequence of bytes. Hence it is legal to append a string to a slice of bytes. Below is the program for that. Notice the ‘…’ at then end of the string
}
