package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetName(name string) *Person {
	p.name = name
	return p
}

func (p *Person) SetAge(age int) *Person {
	p.age = age
	return p
}

func (p *Person) Print() *Person {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
	return p
}

func main() {
	person := &Person{}
	person.SetName("Alice").SetAge(25).Print()
}
