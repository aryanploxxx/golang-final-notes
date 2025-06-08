package main

import (
	"context"
	"time"
)

// Context in go is used for controlling timeouts, cancelling go routines, and passing metadata across your go program
// context can also be used in middlewares to pass data across the request lifecycle
// COntext is like a bucket for your information
// We can assign and store id's to each request

func main() {
	ctx := context.Background()
	exampleTimeout(ctx)
}

func exampleTimeout(ctx context.Context) {
	// ctx := context.Background()
	// this will create an empty struct, which we can use

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	// this will create a new context with a timeout of 2 seconds
	defer cancel()

	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		// simulate an api taking 3 seconds to respond
		close(done)
	}()

	select {
	case <-done:
		println("api call done")
	case <-ctxWithTimeout.Done(): // listen for the context to be done
		println("oops, time limit exceeded, api call timeout", ctxWithTimeout.Err())
		// this will print run because the api call took more than 2 seconds (out timeout limit)
	}
}
