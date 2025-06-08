package main

import "fmt"

func main() {
	src := []int{1, 2, 3, 4, 5}
	numberOfElementsCopied := copy(src, src[3:])

	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("src: %v\n", src)
}

// Number Of Elements Copied: 2
// src: [4 5 3 4 5]
