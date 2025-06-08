package main

import "fmt"

type area func(int, int) int

func main() {
	var areaF area = func(a, b int) int {
		return a * b
	}
	print(2, 3, areaF)
}

func print(x, y int, a area) {
	fmt.Printf("Area is: %d\n", a(x, y))
}

// In this example also we create a user-defined function type area. Then we create a function getAreaFunc() which returns the function of type area

// func sum_avg(a, b int) (int, int)
// As a convention error is returned as the last argument in a function. Example
// func sum(a, b int) (int, error)
