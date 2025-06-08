package main

import "fmt"

func main() {
	count := 0
	for i := 1; i <= 5; i++ {
		func() {
			count++
			fmt.Println(count)
		}()
	}
}

// In the below example, the closure function is able to access the count variable, as well as the value of the count variable, which is retained between different function calls.

// 1
// 2
// 3
// 4
// 5
