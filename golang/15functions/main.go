package main

import "fmt"

func main() {
	fmt.Println("Functions in Go")
	greeter()
	// We can call the function from anywhere in the program

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro Adder Result is: ", proResult)

	proResult2, myMessage := proAdder2(1, 2, 3, 4, 5)
	fmt.Println("Pro Adder Result is: ", proResult2)
	fmt.Println("Pro Adder Message is: ", myMessage)

	// Just writing greeter would mean that we passing the reference of the function
	/*
		func greeterTwo() {
			fmt.Println("Another function")
		}
		greeterTwo()
	*/
	// We cannot define a fucntion in a function in Go
}

// We created main function but we did not call it, still when we run the program, it gets executed
// This is because main function is the entry point of the program

func greeter() {
	fmt.Println("Hello World!")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// to add n arguements when the value of n is not known
func proAdder(values ...int) int {
	total := 0
	// we will need a for loop to add because 'values' is now a slice
	for _, val := range values {
		total += val
	}
	return total
}

// ... indicate Variadic Functions - Variadic functions can be called with any number of trailing arguments,  allow you to pass a variable number of arguments to a function.
// We need to specify what type of value to pass on and what type of value the function will return

// We can return multiple values from a function
func proAdder2(values ...int) (int, string) {
	total := 0
	// we will need a for loop to add because 'values' is now a slice
	for _, val := range values {
		total += val
	}
	return total, "Hi from proAdder2"
}

// func () {
// 	fmt.Println("IIFE!")
// }()

// IIFE - Immediately Invoked Function Expression and Lambda functions also exist in Go
