package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache" // Import the memcache package
)

func main() {
	// Create a new memcache client with the given server addresses
	mC := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11211", "10.0.0.4:11211")

	// Set a new item in the memcache with key "data" and value "30"
	mC.Set(&memcache.Item{Key: "data", Value: []byte("30")})

	// Retrieve the item with key "data" from the memcache
	dataItem, err := mC.Get("data")
	fmt.Println("Data:", string(dataItem.Value))

	// Check if there was an error retrieving the item
	if err != nil {
		// Handle the error
	}

	// Use the retrieved item (dataItem) as needed
}
