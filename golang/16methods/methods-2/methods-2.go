package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func (e employee) printName() employee {
	fmt.Printf("Name: %s\n", e.name)
	return e
}

func (e employee) printAge() employee {
	fmt.Printf("Age: %d\n", e.age)
	return e
}

func (e employee) printSalary() {
	fmt.Printf("Salary: %d\n", e.salary)
}

// Method chaining is a technique to call multiple methods one after the other on the same object. It is achieved by returning the object from the method.
// Method chaining in Go refers to the practice of calling multiple methods on the same object or value in a single statement, where each method call returns the object itself or a modified version of it. This pattern allows for more concise and readable code by chaining method calls together.

func main() {
	emp := employee{name: "Sam", age: 31, salary: 2000}
	emp.printName().printAge().printSalary()
}
