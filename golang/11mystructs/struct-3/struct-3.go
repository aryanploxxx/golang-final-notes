package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// struct-2 and struct-3 are similar

func main() {
	p := &Person{
		Name: "Alice",
		Age:  30,
	}

	changeDetails(p)

	fmt.Println("Changed Person name:", p.Name)
	fmt.Println("Changed Person age:", p.Age)
}

func changeDetails(p *Person) {
	p.Name = "Bob"
	p.Age = 25
}

/*
	The & operator can be used to get the pointer to a struct variable.
		emp := employee{name: "Sam", age: 31, salary: 2000}
		empP := &emp
	struct pointer can also be directly created as well
		empP := &employee{name: "Sam", age: 31, salary: 2000}

	O/P:
		Emp: &{name:Sam age:31 salary:2000}
		Emp: &{name:John age:30 salary:3000}
*/

/*
	There are two ways of creating a pointer to the struct
	- Using the & operator
	- Using the new keyword
*/
