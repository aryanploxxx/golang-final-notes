package main

import "fmt"

// A struct can have anonymous fields as well, meaning a field having no name. The type will become the field name. In below example, string will be the field name as well

type employee struct {
	string
	age    int
	salary int
}

func main() {
	emp := employee{string: "Sam", age: 31, salary: 2000}
	//Accessing a struct field
	n := emp.string
	fmt.Printf("Current name is: %s\n", n)
	//Assigning a new value
	emp.string = "John"
	fmt.Printf("New name is: %s\n", emp.string)

}
