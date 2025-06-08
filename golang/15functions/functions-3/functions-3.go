package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))
}

func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// If you already have a slice and you need to pass it as a variadic param then that can be done by adding three dots (…) after the argument while calling the function

// var numbers := []int{2,3,5}
// add(numbers...)
