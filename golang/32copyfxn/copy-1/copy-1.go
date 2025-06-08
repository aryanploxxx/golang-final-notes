package main

import "fmt"

// func copy(dst, src []Type) int

func main() {
	src := "abc"
	dst := make([]byte, 3)

	numberOfElementsCopied := copy(dst, src)
	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)
}

// Number Of Elements Copied: 3
// dst: [97 98 99]
// src: abc
// -> A string in go is nothing but a sequence of bytes. Hence it is legal to copy a string to a slice of bytes.
