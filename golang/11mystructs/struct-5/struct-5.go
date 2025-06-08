package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	emp1 := employee{}
	fmt.Printf("Emp1: %+v\n", emp1)

	emp2 := employee{name: "Sam", age: 31, salary: 2000}
	fmt.Printf("Emp2: %+v\n", emp2)

	emp3 := employee{
		name:   "Sam",
		age:    31,
		salary: 2000,
	}
	fmt.Printf("Emp3: %+v\n", emp3)

	emp4 := employee{
		name: "Sam",
		age:  31,
	}
	fmt.Printf("Emp4: %+v\n", emp4)
}

// Emp1: {name: age:0 salary:0}
// Emp2: {name:Sam age:31 salary:2000}
// Emp3: {name:Sam age:31 salary:2000}
// Emp4: {name:Sam age:31 salary:0}

// It is also ok to initialize only some of the fields with value. The field which are not initialized with value will get the default zero value of their type

// struct can also be initialized without specifying the field names. But in this case, all values for each of the field has to be provided in sequence
// emp := employee{"Sam", 31, 2000}
