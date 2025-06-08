package main

import (
	"fmt"     // For printing
	"strings" // For string manipulation
	// "sync" For synchronization (though WaitGroup isn't used in this code)
)

// Global variables declaration
var initialString string // Input string
var initialBytes []byte  // Byte slice of input string
var stringLength int     // Length of input string
var finalString string   // Result string
var lettersProcessed int // Counter for processed letters
// var wg sync.WaitGroup      // WaitGroup (unused in this code)
var applicationStatus bool // Flag to control program termination

// Function to send letters to channel
func getLetters(gQ chan string) {
	// Iterate through each byte in initialBytes
	for i := range initialBytes {
		gQ <- string(initialBytes[i]) // Send each letter to channel
	}
}

// Function to capitalize letters received from channel
func capitalizeLetters(gQ chan string) {
	for {
		// Check if all letters are processed
		if lettersProcessed >= stringLength {
			applicationStatus = false // Set flag to end program
			break
		}
		select {
		case letter := <-gQ: // Receive letter from channel
			capitalLetter := strings.ToUpper(letter) // Capitalize letter
			finalString += capitalLetter             // Add to final string
			lettersProcessed++                       // Increment counter
		}
	}
}

func main() {
	// Initialize program status
	applicationStatus = true

	// Create channels
	getQueue := make(chan string)   // Channel for passing letters
	stackQueue := make(chan string) // Unused channel

	// Initialize input string with Gettysburg Address excerpt
	initialString = "Four score and seven years ago..."
	initialBytes = []byte(initialString) // Convert to byte slice
	stringLength = len(initialString)    // Get string length
	lettersProcessed = 0                 // Initialize counter

	fmt.Println("Let's start capitalizing")

	// Start concurrent operations
	go getLetters(getQueue)     // Start sender goroutine
	capitalizeLetters(getQueue) // Start capitalizing (in main goroutine)

	// Close channels (though this happens after capitalizeLetters completes)
	close(getQueue)
	close(stackQueue)

	// Wait for completion
	for {
		if !applicationStatus {
			fmt.Println("Done")
			fmt.Println(finalString)
			break
		}
	}
}
