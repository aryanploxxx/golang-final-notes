package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Initialize a logger
var panicLogger *log.Logger

// Centralized recovery function to log panics
func handlePanic() {
	if rec := recover(); rec != nil {
		// Log the panic details
		panicLogger.Printf("Recovered from panic: %v at %v\n", rec, time.Now().Format(time.RFC3339))
		fmt.Println("recovered", rec)
	}
}

// Function to simulate file operations
func getFileDetails(fileName string) {
	defer handlePanic() // Defer recovery

	fmt.Println("Fetching file details...")
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error accessing file details: %v", err)) // Panic on error
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("File Size:", fileInfo.Size())
}

// Function to open a file
func openFile(fileName string) {
	defer handlePanic() // Defer recovery

	fmt.Println("Opening file...")
	if _, err := os.Stat(fileName); err != nil {
		panic(fmt.Sprintf("Error: File '%s' does not exist", fileName)) // Panic on error
	}
	fmt.Println("File opened successfully.")
}

// Main function
func main() {
	// Set up the logger to write to a file
	logFile, err := os.OpenFile("panic_logs.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		// fmt.Println("Failed to open log file:", err)
		// err:="guru"
		defer handlePanic()
		panic(fmt.Sprintf("Panic occured : %v", err))

		// return
	}
	defer logFile.Close()
	panicLogger = log.New(logFile, "PANIC_LOG: ", log.LstdFlags)

	fmt.Println("Enter file name:")
	var fileName string
	fmt.Scanln(&fileName) // Get input from the user

	// Perform file operations
	openFile(fileName)
	getFileDetails(fileName)

	fmt.Println("Program executed successfully.")
}
