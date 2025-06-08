package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	go add(2, 3) // ❌ Return value is ignored
	fmt.Println("Main function continues...")
}

// Some points to note about nil channel
// 	Sending to a  nil channel blocks forever
// 	Receiving from nil channel blocks forever
// 	Closing a nil channel results in panic
