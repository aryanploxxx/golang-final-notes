package main

import "fmt"

type Address struct {
	city    string
	country string
}

type User struct {
	title   string
	salary  int
	address Address
}

func (us *User) changeName(name string) {
	us.title = name
}

func (u *User) changeAddress() {
	u.address.city = "Mumbai"
	u.address.country = "India"
}

func main() {

	var x User
	fmt.Printf("%T", x)
	// main.Person

	y := User{}
	fmt.Printf("%T", y)
	// main.Person

	p := new(User)
	fmt.Printf("%T", p)
	// *main.Person
	fmt.Printf("%p", p)
	// 0xc00011e000 -> pointer address

	/*
		p := Person{"Alice", 30, "alice@example.com"}
		p := Person{
			Name:  "Alice",
			Age:   30,
			Email: "alice@example.com",
		}
	*/

	/*
		Using the  new() keyword will:
		- Create the struct
		- Initialize all the field to the zero default value of their type
		- Return the pointer to the newly created struct
	*/

	addressUser := Address{
		city:    "New Delhi",
		country: "India",
	}

	user1 := User{
		title:   "Software Developer",
		salary:  10000,
		address: addressUser,
	}

	user1.changeName("aryan")
	user1.changeAddress()

	fmt.Println("user", user1)
	fmt.Println("user address", addressUser)
}
