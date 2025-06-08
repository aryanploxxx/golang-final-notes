package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var value int = 10
	fmt.Println("Value is:", value)
	fmt.Println("Address of value is:", &value)
	/*
		var ptr *int
		fmt.Println("Value in pointer is:", ptr)
		// Value in pointer is: <nil>
		// Initially the value will be nil
	*/

	myNumber := 23
	// Now we want to create a pointer to this variable
	var newPtr = &myNumber
	// We are referencing the address of myNumber to newPtr
	fmt.Println("Value in new pointer is (address):", newPtr)
	fmt.Println("Value in new pointer is (actual value):", *newPtr)
	// Value in new pointer is: 0xc00000a100

	*newPtr = *newPtr + 2
	fmt.Println("Value in pointer is:", myNumber)
}
