package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	currentTime := time.Now().UTC()

	// Convert time.Time to string in RFC3339 format
	formattedTime := currentTime.Format(time.RFC3339)
	fmt.Println("Formatted Time:", formattedTime)
}
