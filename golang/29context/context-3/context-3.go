package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Every HTTP REQUEST has a built in context into them, which is a struct that holds the request data

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	// We are waiting only 2 seconds for the API Response
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		// API Response is simulatrd to take 3 seconds
		fmt.Println("API Response")
	case <-ctx.Done(): // we are listening for the context to be done in the channel
		fmt.Println("API Timeout; Context Expired")
		http.Error(w, "Request Context Timeout", http.StatusRequestTimeout)
		return
	}
}
