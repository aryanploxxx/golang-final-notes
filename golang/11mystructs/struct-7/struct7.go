package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary int
}

func main() {
	empP := new(employee)
	fmt.Printf("Emp Pointer Address: %p\n", empP)
	fmt.Printf("Emp Pointer: %+v\n", empP)
	fmt.Printf("Emp Value: %+v\n", *empP)

	x := employee{name: "Sam", age: 31, salary: 2000}
	fmt.Printf("Emp: %v\n", x)
	fmt.Printf("Emp: %+v\n", x)
	fmt.Printf("Emp: %#v\n", x)
	fmt.Println(x)

}

// Emp Pointer Address: 0xc000130000
// Emp Pointer: &{name: age:0 salary:0}
// Emp Value: {name: age:0 salary:0}

// Emp: {Sam 31 2000}
// Emp: {name:Sam age:31 salary:2000}
// Emp: main.employee{name:"Sam", age:31, salary:2000}
// {Sam 31 2000}
