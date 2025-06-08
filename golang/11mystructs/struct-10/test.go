package main

import "fmt"

// A struct can have another struct nested in it. Let’s see an example of a nested struct. In below employee struct has address struct nested it in.

func check() {
	address := address{City: "London", Country: "UK"}
	emp := employee{Name: "Sam", Age: 31, Salary: 2000, Address: address}
	fmt.Printf("City: %s\n", emp.Address.City)
	fmt.Printf("Country: %s\n", emp.Address.Country)
}
