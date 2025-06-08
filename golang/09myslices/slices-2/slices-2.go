package main

import "fmt"

func main() {
	numbers := make([]int, 3, 3)
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Append number 4
	numbers = append(numbers, 4)
	fmt.Println("\nAppend Number 4")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

// numbers=[1 2 3]
// length=3
// capacity=3

// Append Number 4
// numbers=[1 2 3 4]
// length=4
// capacity=6

// When slice length is equal than capacity.
// In this case since there is no more capacity, so no new elements can be accommodated.
// So in this case under the hood an array of double the capacity will be allocated.
// The current array pointed by the  slice will be copied to that new array.
// Now the slice will starting pointing to this new array.
// Hence the capacity will be doubled and length will be increased by 1.
