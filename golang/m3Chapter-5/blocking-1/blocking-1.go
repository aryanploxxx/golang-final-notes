package main

import (
	"fmt"
	"time"
)

// func thinkAboutKeys() {
// 	for {
// 		fmt.Println("Still Thinking")
// 		time.Sleep(1 * time.Second)
// 	}
// }

func thinkAboutKeys(bC chan int) {
	i := 0
	max := 10
	for {
		if i >= max {
			bC <- 1
		}
		fmt.Println("Still Thinking")
		time.Sleep(1 * time.Second)
		i++
	}
}

func main() {
	fmt.Println("Where did I leave my keys?")

	blockChannel := make(chan int)
	// In Go, channels are reference types. When you pass a channel to a function, you're passing a reference to the same underlying channel, not a copy of the channel itself.
	// Channels in Go are designed for synchronization and communication between goroutines. If channels were passed by value (i.e., if copies were made), each copy would have its own independent state, breaking the intended behavior of communication.
	go thinkAboutKeys(blockChannel)
	<-blockChannel

	fmt.Println("OK I found them!")
}
