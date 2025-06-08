package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	// Channels are a way or a pipeline where multiple goroutines interact, hence we need to specify what kind of values, we will be passing through them
	/*
		// Creating a channel
		myCh := make(chan int)

		// Sending a value to the channel
		myCh <- 5

		// Printing the value received from the channel
		fmt.Println(<-myCh)

		-> The above code will result in error, because channels only allow inserting a value in them if there is a goroutine waiting to receive the value from the channel
		-> Running the above code will result in a deadlock, as the channel is not closed and the goroutine is waiting for the value to be received from the channel
	*/

	myCh := make(chan int, 2)
	/*
		myCh := make(chan int)
		-> This is an unbuffered channel, which means that the channel will only accept a value if there is a goroutine waiting to receive the value from the channel
		-> if there were 2 values being inserted, then we would need 2 goroutines to receive the values from the channel
		-> If there is no goroutine waiting to receive the value from the channel, then the program will result in a deadlock

		myCh := make(chan int, 2)
		-> This is a buffered channel, which means that the channel will accept the values even if there is no goroutine waiting to receive the value from the channel
		-> The values will be stored in the buffer, and the goroutine can receive the values from the buffer
		-> if we insert more values than the buffer size, then it will just simple ignore the values without any error
	*/
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// <-
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// ch <-chan int -- this symbol makes it a RECIEVE Only channel -- we won't even be allowed to close a channel in this way
		fmt.Println("Goroutine 2")

		value, isChannelOpen := <-ch
		fmt.Println(isChannelOpen) // return true, false based on which we can determine if value returned by <-ch is actual value or value due to listening on closed channel
		fmt.Println(value)
		wg.Done()
	}(myCh, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// ch chan<- int -- this symbol makes it a SEND Only channel
		fmt.Println("Goroutine 1")
		// close(ch) // if we close here (at this line) , we would encounter an error that we are listening to a closed channel
		ch <- 5
		// By default if we listen to a closed channel, it will return the zero value (0) of the type of the channel
		// this would mean listening on closed channel and listening on channel with ch <- 0 would result in same output -> which is problematic
		// -> to solve this problem we can use a second return value from the channel, which will tell us if the channel is closed or not, done in goroutine 2

		close(ch) // Closing the channel
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
