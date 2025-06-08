package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

// add, done, wait are the methods that we can use to make the main function wait for the goroutines to finish
// waitgroup is an advanced versoion of sleep() method
// Add() -> as soon as a goroutine is created, add that to the waitgroup
// waitgroup will not end main function until all the goroutines are finished
// Done() -> as soon as the goroutine is finished, call done() method. it is our responsibility to call done() method
// wg.Wait() -> this will make the main function wait for all the goroutines to finish, almost always goes to the end of the main function or the function it is being called inside of

func main() {
	// go greeter("Hello")
	// greeter("World")
	// if we run only this much without writing sleep() in greeter, then only world will be printed 5 times
	// this happens because the main function is not waiting for the goroutine to finish and finishes as soon as the world is printed 5 times
	// adding sleep() in greeter will make the main function wait for the goroutine to finish and the output would be something like Hwllo and World printed in some order

	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
	}
	// this will take some time to execute, to make them fast, we can execute them in seperate threads
	/*
		for _, website := range websiteList {
			getStatusCode(website)
		}
		-> this will work just fine, but it will take some time to execute as every request is made one after the other
		-> we can make them in parallel using goroutines
		-> by just adding go keyword before the function call, we can make them run in parallel
		-> however running them using 'go' keyword will not return any output as the main function will not wait for the goroutines to finish
		-> one workaround is to use sleep method in the function
		-> however, this is not a good practice as we are making the main function wait for the goroutines to finish
		->> we'll use the 'sync' package to make the main function wait for the goroutines to finish
	*/

	for _, website := range websiteList {
		// getStatusCode(website)
		go getStatusCode(website)
		// this will make the getStatusCode function run in a seperate threads and execute much faster, however, we are not waiting in the main method hence it will not return any output
		//  one workaround is to use sleep method in the function
		wg.Add(1)
	}

	wg.Wait() // this will make the main function wait for all the goroutines to finish
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// Some of our APIs talk to the read servers only, some talk to read-write servers both
// We can use goroutines to make the read-write API calls in parallel
// we can launch threads using goroutines to talk to different microservices in parallel

func getStatusCode(endpoint string) {
	defer wg.Done()
	// this signal will pass every time a function will be finished so that wg.Wait() can check how many goroutines are still running

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint) // this is the memory where all the goroutines are writing to, hence we need to lock it
		mut.Unlock()

		fmt.Printf("%d status code for %s", result.StatusCode, endpoint)
	}

}

/*
var wg sync.WaitGroup
go addToFinalStack(letterChannel, &wg)

var wg *sync.WaitGroup
go addToFinalStack(letterChannel, wg)


*/
