package main

import "fmt"

func main() {
	fmt.Println("Structs")
	// There is no inheritance in Go, no super parent concept

	aryan := User{"Aryan", "aryan@mail.com", true, 21}

	fmt.Println("Person: ", aryan)
	// Person:  {Aryan aryan@mail.com true 21}

	fmt.Printf("Aryan details are: %+v\n", aryan)
	// Aryan details are: {Name:Aryan Email:aryan@mail.com Status:true Age:21}
	// fmt.Printf("Aryan details are: %+v\n", aryan) -> %v will give same result as println
	// %+v will print the field names along with the values

	fmt.Printf("Name is %v and Email is %v.\n", aryan.Name, aryan.Email)
	// Name is Aryan and Email is aryan@mail.com.
}

// Created Globally
// All below variables start with capital letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Structs - Group related data together in one datatype
/*
	type Employee struct {
		Name 	 string
		Age 	 int
		isRemote bool
	}

	func main() {
		employee1 := Employee{
			name: "Aryan",
			age: 21,
			isRemote: true,
		}
	}

	fmt.Println("Employee Name", employee1.Name)
	fmt.Println("Employee Name", employee1.Age)
*/

// We can define Anonymous Structs - More like one time struct to reduce overhead inside the functions
/*
	jobs := struct {
		title  string
		salary int
	}{
		title:  "Software Developer",
		salary: 10000,
	}
	fmt.Println("Job Type", job.title)
*/
