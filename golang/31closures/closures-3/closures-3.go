package main

import "fmt"

func main() {
	valueOutside := "somevalue"
	func() {
		fmt.Println(valueOutside)
	}()
}

// somevalue
// Below is also another example of a closure function. The function is able to access the valueOutside variable.
