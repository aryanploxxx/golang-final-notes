package main

import "fmt"

func main() {
	numbers1 := []int{1, 2}
	numbers2 := []int{3, 4}
	numbers := append(numbers1, numbers2...)
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

// It is also possible to append one slice to another slice. Below is the format for that.
// res := append(slice1, slice2...)
// Notice ‘…’ after the second slice. ‘…’ is the operator which means that the argument is a variadic parameter. Meaning that during run time slice2 will be expanded to its individual elements which are passed as multiple arguments to the append function.
