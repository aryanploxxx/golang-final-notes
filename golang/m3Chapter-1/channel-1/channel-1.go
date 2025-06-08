// Package declaration
package main

// Import necessary packages
import (
	"fmt"     // For printing
	"runtime" // For controlling runtime behavior
	"strings" // For string manipulation
	"sync"    // For synchronization primitives
)

// Global variables declaration
var initialString string // Will hold the input string
var finalString string   // Will hold the transformed string
var stringLength int     // Will store the length of input string

// Function that receives a letter from channel and adds it to final string
func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel // Receive letter from channel
	finalString += letter     // Append to final string
	wg.Done()                 // Signal completion to WaitGroup
}

// Function that capitalizes a letter and sends it to channel
func capitalize(letterChannel chan string, currentLetter string, wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter) // Convert letter to uppercase
	letterChannel <- thisLetter                  // Send capitalized letter to channel
	wg.Done()                                    // Signal completion to WaitGroup
}

func main() {
	// Set maximum number of CPUs that can execute simultaneously to 2
	runtime.GOMAXPROCS(2)

	// Create WaitGroup for synchronization
	var wg sync.WaitGroup

	// Initialize input string with Gettysburg Address excerpt
	initialString = "Four score and seven years ago..."

	// Convert string to byte slice for character-by-character processing
	initialBytes := []byte(initialString)
	stringLength = len(initialBytes)

	// Create unbuffered channel for passing letters between goroutines
	var letterChannel chan string = make(chan string)

	// Process each character in the string
	for i := 0; i < stringLength; i++ {
		wg.Add(2) // Add 2 to WaitGroup counter for two goroutines

		// Launch two goroutines:
		// 1. capitalize the current letter
		go capitalize(letterChannel, string(initialBytes[i]), &wg)
		// 2. receive the capitalized letter and add to final string
		go addToFinalStack(letterChannel, &wg)

		// Wait for both goroutines to complete before next iteration
		wg.Wait()
	}

	// Print the final capitalized string
	fmt.Println(finalString)
}
