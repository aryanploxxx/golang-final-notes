package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Person struct {
	Name    string
	Age     int
	Email   string
	Address // Embedding Address struct -> Anonymous nested struct fields
	// Address Address -> if defined this way, we will have to access fields via this way - alice.Address.City
}

func main() {
	alice := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
		Address: Address{
			Street: "123 Main St",
			City:   "Wonderland",
			State:  "Fantasy",
			Zip:    "12345",
		},
	}

	// Accessing promoted fields directly
	fmt.Println("Name:", alice.Name)         // Output: Name: Alice
	fmt.Println("Name:", alice.Age)          // Output: Age: 30
	fmt.Println("Name:", alice.Email)        // Output: Email: alice@example.com
	fmt.Println("Street:", alice.Street)     // Output: Street: 123 Main St
	fmt.Println("City:", alice.Address.City) // Output: City: Wonderland
	fmt.Println("State:", alice.State)       // Output: State: Fantasy
	fmt.Println("Zip:", alice.Zip)           // Output: Zip: 12345
}
