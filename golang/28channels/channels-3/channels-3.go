package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sum(ch, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	time.Sleep(time.Second * 1)
}

func sum(ch chan int, len int) {
	sum := 0
	for i := 0; i < len; i++ {
		sum += <-ch
	}
	fmt.Printf("Sum: %d\n", sum)
}

// Sending on a close channel will cause a panic.
// Also closing a already closed channel will cause a panic

/*
While receiving from a  channel we can also use an additional variable to determine if the channel has been closed.  Below is the syntax for the  same
val,ok <- ch
The value of ok will be

True if the channel is not closed
False if the channel is closed

*/
