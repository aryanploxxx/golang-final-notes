package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// struct-2 and struct-3 are similar

func main() {
	p := Person{
		Name: "Alice",
		Age:  30,
	}

	changeDetails(&p)

	fmt.Println("Changed Person name:", p.Name)
	fmt.Println("Changed Person age:", p.Age)
}

func changeDetails(p *Person) {
	p.Name = "Bob"
	p.Age = 25
}
