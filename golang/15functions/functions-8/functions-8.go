package main

import "fmt"

type area func(int, int) int

func main() {
	areaF := getAreaFunc()
	print(2, 3, areaF)

}

func print(x, y int, a area) {
	fmt.Printf("Area is: %d\n", a(x, y))
}

func getAreaFunc() area {
	return func(x, y int) int {
		return x * y
	}
}
