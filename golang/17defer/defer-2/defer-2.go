package main

import "fmt"

func main() {
	sample := "abc"

	defer fmt.Printf("In defer sample is: %s\n", sample)
	sample = "xyz"
}

// In defer sample is: abc
/*
	Evaluation of defer arguments
	defer arguments are evaluated at the time defer statement is evaluated
*/

// In case we have multiple defer functions within a particular function, then all the  defer functions will be executed in last in first out order
/*
	package main
	import "fmt"
	func main() {
		i := 0
		i = 1
		defer fmt.Println(i)
		i = 2
		defer fmt.Println(i)
		i = 3
		defer fmt.Println(i)
	}
*/
