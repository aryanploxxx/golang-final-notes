package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go goOne(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

func goOne(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "From goOne goroutine"
}

// goOne will take 2 seconds will complete, but before that only, after 1 second,
// time.After(time.Second * 1) will send a value to the channel making one of the case execute in select, and the fxn. will finish
