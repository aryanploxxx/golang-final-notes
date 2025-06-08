package main

import "fmt"

func main() {
	handle(1, "abc")
	handle("abc", "xyz", 3)
	handle(1, 2, 3, 4)
}

func handle(params ...interface{}) {
	fmt.Println("Handle func called with parameters:")
	for _, param := range params {
		fmt.Printf("%v\n", param)
	}
}
