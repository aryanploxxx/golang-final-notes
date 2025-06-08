package main

import (
	"fmt"     // For printing
	"strings" // For string operations
)

// Function that returns another function (closure)
func shortenString(message string) func() string {
	// Returns an anonymous function that has access to the 'message' parameter
	return func() string {
		// Split the message into slice of words
		messageSlice := strings.Split(message, " ")

		// Get number of words in slice
		wordLength := len(messageSlice)
		fmt.Println(wordLength)

		// Check if there are any words left
		if wordLength < 1 {
			return "Nothingn Left!" // Note: there's a typo in "Nothing"
		} else {
			// Remove the last word by slicing up to (but not including) the last element
			messageSlice = messageSlice[:(wordLength - 1)]

			// Join the remaining words back into a string with spaces
			message = strings.Join(messageSlice, " ")

			// Print and return the shortened message
			fmt.Println(message)
			return message
		}
	}
}

func main() {
	// Create a closure with an empty string
	myString := shortenString("Welcome to concurrency in Go! ...")

	// Call the closure multiple times
	(myString())
	(myString())
	(myString())
	(myString())
	(myString())
	(myString())
}
