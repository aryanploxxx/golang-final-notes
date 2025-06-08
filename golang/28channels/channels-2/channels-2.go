package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	fmt.Printf("Capacity: %d\n", cap(ch))

	ch <- 5
	fmt.Printf("Len: %d\n", len(ch))

	ch <- 6
	fmt.Printf("Len: %d\n", len(ch))

	ch <- 7
	fmt.Printf("Len: %d\n", len(ch))
}

// Len: 1
// Len: 2
// Len: 3
