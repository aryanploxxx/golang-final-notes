package main

import (
	"fmt"
)

func main() {
	modulus := getModulus()
	modulus(-1)
	modulus(2)
	modulus(-5)
}

func getModulus() func(int) int {
	count := 0
	return func(x int) int {
		count = count + 1
		fmt.Printf("modulus function called %d times\n", count)
		if x < 0 {
			x = x * -1
		}
		return x
	}
}

// modulus function called 1 times
// modulus function called 2 times
// modulus function called 3 times

// The getModulus function returns a closure. It is assigned to a variable modulus
// This closure function can access the count variable defined outside its body.
// The value of the count variable is retained between different function calls of modulus function
