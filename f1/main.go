package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records (rows) from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Check if the file contains any data
	if len(records) == 0 {
		fmt.Println("No data found in the file.")
		return
	}

	// Print header (optional)
	header := records[0]
	fmt.Println("Header:", header)

	// Iterate through the rows (skipping the header)
	for i, row := range records[1:] {
		if len(row) != len(header) {
			fmt.Printf("Skipping malformed row %d: %v\n", i+1, row)
			continue
		}

		// Example: Print row details
		fmt.Printf("Row %d: Name=%s, Age=%s, City=%s\n", i+1, row[0], row[1], row[2])
	}
}
