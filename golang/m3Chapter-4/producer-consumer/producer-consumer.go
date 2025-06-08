package main

import (
	"fmt"
)

var comm = make(chan bool)
var done = make(chan bool)

func producer() {
	for i := 0; i < 10; i++ {
		comm <- true
	}
	done <- true
}

func consumer() {
	for {
		communication := <-comm
		fmt.Println("Communication from producer received!", communication)
	}
}

func main() {
	go producer()
	go consumer()
	<-done
	// The program then blocks on <-done, waiting for the producer to send a signal on the done channel.
	// Once the producer signals that it is done, the program prints "All Done!" and exits.

	fmt.Println("All Done!")
}
