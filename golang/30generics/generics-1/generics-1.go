package main

import "fmt"

// #1
// type stack struct {
// 	elements []int
// }

type stack[T any] struct {
	elements []T
}

func main() {
	// #1
	// myStack := stack{
	// 	elements: []int{1, 2, 3, 4, 5},
	// }
	// fmt.Println("Stack elements:", myStack)
	// Stack elements: {[1 2 3 4 5]}

	myStack := stack[string]{
		elements: []string{"golang"},
	}
	fmt.Println("Stack elements:", myStack)
}
