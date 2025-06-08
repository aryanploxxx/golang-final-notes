package main

import "fmt"

func main() {
	i, j := 42, 110
	fmt.Println("i:", i, "\nj:", j)
	fmt.Println("Address of i:", &i, "\nAddress of j:", &j)

	// i: 42
	// j: 110
	// Address of i: 0xc00000a0f8
	// Address of j: 0xc00000a110

	p := &i
	fmt.Println("Value of p:", p)
	fmt.Println("Value of i:", *p)

	// 0xc00000a0f8

	radius := 10
	squared(&radius)
	fmt.Println("Squared value is:", radius)

}

type person struct {
	name string
	age  int
}

func initPerson(p *person) *person {
	m := person{name: "John", age: 25}
	// fmt.Println("initPerson:", &m)
	return &m
}
func squared(radius *int) {
	*radius = *radius * *radius
}
