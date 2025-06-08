package main

import "fmt"

// Return a function from a function in Go (Golang)

func main() {
	areaF := getAreaFunc()
	res := areaF(2, 4)
	fmt.Println(res)
}

func getAreaFunc() func(int, int) int {
	return func(x, y int) int {
		return x * y
	}
}
