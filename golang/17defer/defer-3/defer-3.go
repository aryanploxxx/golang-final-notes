package main

import "fmt"

func main() {
	s := test()
	fmt.Println(s)
}
func test() (size int) {
	defer func() { size = 20 }()
	size = 30
	return
}

// In case of named return value in the function, the defer function can read as well as modified those named return values. If the defer function modifies the name return value then that modified value will  be returned
// O/P: 20
