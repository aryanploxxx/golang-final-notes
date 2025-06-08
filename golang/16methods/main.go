package main

import "fmt"

// A method in golang is nothing but a function with a receiver
// The method will have access to the properties of the receiver and can call the receiver’s other methods
// Since method lets you define a function on a type, it lets you write object-oriented code in Golang. There are also some other benefits such as two different methods can have the same name in the same package which is not possible with functions

// func (receiver receiver_type) some_func_name(arguments) return_values

// There can exist different methods with the same name with a different receiver, but there cannot exist two different functions with the same name in the same package.

func main() {
	fmt.Println("Methods")
	// There is no inheritance in Go, no super parent concept

	aryan := User{"Aryan", "aryan@mail.com", true, 21, 21}

	fmt.Println("Person: ", aryan)
	fmt.Printf("Aryan details are: %+v\n", aryan)
	fmt.Printf("Name is %v and Email is %v.\n", aryan.Name, aryan.Email)

	aryan.GetStatus()
	aryan.NewMail()
	/*
		Person:  {Aryan aryan@mail.com true 21 21}
		Aryan details are: {Name:Aryan Email:aryan@mail.com Status:true Age:21 oneAge:21}
		Name is Aryan and Email is aryan@mail.com.
		Is user active?:  true
		New email is:  changed@mail.com
		-> Interesting thing to note that the original value of the struct email is not changed
	*/
}

// Created Globally
// All below variables start with capital letter
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
	// oneAge is not accessible outside the package as first letter is small
}

/*
	func (u User) GetStatus() {

	}
	-> Now we have converted a function into a method by passing the entire struct into it
*/

func (u User) GetStatus() {
	fmt.Println("Is user active?: ", u.Status)
}

// Note that, here a copy of the struct is passed
func (u User) NewMail() {
	u.Email = "changed@mail.com"
	fmt.Println("New email is: ", u.Email)
}

// func (u *User) NewMail() -> pass this way if you want to pass it by reference; pointer of user type
