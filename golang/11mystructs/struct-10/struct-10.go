package main

type employee struct {
	Name    string
	Age     int
	Salary  int
	Address address
}

type address struct {
	City    string
	Country string
}

// A struct can have another struct nested in it. Let’s see an example of a nested struct. In below employee struct has address struct nested it in.

func main() {
	check()
}
