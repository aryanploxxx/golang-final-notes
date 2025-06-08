package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p := Person{
		Name:  "Bob",
		Age:   25,
		Email: "bob@example.com",
	}

	op := p.Greet()
	fmt.Println("Greet: ", op)
	p.UpdateEmail("newmail@example.com")
	fmt.Println("Updated Email: ", p.Email)

}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func (p *Person) UpdateEmail(newEmail string) {
	p.Email = newEmail
}

/*
	Value vs. Pointer Receivers
	When defining methods on structs, you can choose between value receivers and pointer receivers. The choice depends on whether you need to modify the receiver or not.

	-> Value Receivers: Use value receivers when the method does not modify the receiver’s fields. Value receivers are a copy of the original struct.
	func (p Person) Greet() string {
		return "Hello, my name is " + p.Name
	}

	-> Pointer Receivers: Use pointer receivers when the method needs to modify the receiver’s fields or to avoid copying large structs.
	func (p *Person) UpdateEmail(newEmail string) {
		p.Email = newEmail
	}
*/
