package main

import "fmt"

// A function that sends a simple function to the functionChannel
func sendFunctionToChannel(functionChannel chan func() string) {
	functionChannel <- func() string { return "Message Sent!" }
}

// This code demonstrates how to send and execute a function through a channel in Go
func main() {
	// Create a channel to hold functions that return a string
	functionChannel := make(chan func() string)
	defer close(functionChannel) // Ensure the channel is closed when the program exits

	// Start a goroutine to send a function into the channel
	go sendFunctionToChannel(functionChannel)

	// Use a select statement to receive the function from the channel
	select {
	case receivedFunction := <-functionChannel: // Receive the function from the channel, we just recived the reference to the function
		message := receivedFunction() // Execute the function to get the message
		fmt.Println(message)          // Print the message returned by the function
		fmt.Println("Function Received and Executed!")
	}
}
