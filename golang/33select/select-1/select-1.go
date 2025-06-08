package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	ch <- "From goTwo goroutine"
}

// In the above program, we created two channels that are passed to two different goroutines. Then each of the goroutines is sending one value to the channel. In the select, we have two case statements. Each of the two case statements is waiting for a receive operation to complete on one of the channels. Once any receive operation is complete on any of the channels it is executed and select exits. So as seen from the output, in the above program, it prints the received value from one of the channels and exits.

// The select statement is useful if there are multiple goroutines that are sending data to multiple channels concurrently. The select statement can then receive the data concurrently from any of one goroutine and execute the statement which is ready. So select along with channels and goroutines become a very powerful tool for managing synchronization and concurrency.

// default case will be executed if no send it or receive operation is ready on any of the case statements
// default statement prevents the select from blocking forever
// If the select statement doesn't contain a default case then it can block forever until one send or receive operation is ready on any case statement.
