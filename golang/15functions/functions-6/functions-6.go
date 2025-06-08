package main

import "fmt"

// Higher Order Functions

func main() {
	add, subtract := getAddSubtract()
	print(3, 4, add, subtract)
}

func print(x, y int, add func(int, int) int, subtract func(int, int) int) {
	fmt.Printf("Sum is: %d\n", add(x, y))
	fmt.Printf("Difference Value is: %d\n", subtract(x, y))
}

func getAddSubtract() (func(int, int) int, func(int, int) int) {
	add := func(x, y int) int {
		return x + y
	}
	subtract := func(x, y int) int {
		return x - y
	}
	return add, subtract
}
