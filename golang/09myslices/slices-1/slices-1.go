package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	// -> type, size, capacity
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

	//Append number 5
	numbers = append(numbers, 4)
	fmt.Println("\nAppend Number 5")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}
