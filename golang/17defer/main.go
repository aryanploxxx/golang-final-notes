package main

import "fmt"

// Defer as the name suggests is used to defer the cleanup activities in a function. These cleanup activities will be performed at the end of the function.Execution of a deferred function is delayed to the moment the surrounding function returns
// Execution of a deferred function is delayed to the moment the surrounding function returns
// deferred function will also be executed if the enclosing function terminates abruptly. For example in case of a panic

// We already learn above that defer function is the only function that is called after the panic. So it makes sense to put the recover function in the defer function only. If the recover function is not within the defer function then it will not stop panic.

func main() {
	fmt.Println("Hello 1")
	defer fmt.Println("World 1")
	// Hello World - no significant impact

	defer fmt.Println("Hello 2")
	fmt.Println("World 2")
	// World Hello - defer is executed in reverse order

	/*
		Hello 1
		World 2
		Hello 2
		World 1
	*/

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	// Hello Two One World - defer is executed in reverse order

	myDefer()
	// 4 3 2 1 0

	// A defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
