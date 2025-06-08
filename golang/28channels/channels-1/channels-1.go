package main

import "fmt"

func main() {
	var a chan int
	fmt.Println(a)
	// {nil}

	// To define the channel we can use the inbuilt function make.
	a = make(chan int)
	fmt.Println(a)
	// 0xc0000240c0
}

// By default when we create channel with make, it creates a unbuffered channel which essentially means that channel created cannot store any data.
// So any send on a channel is blocked until there is  another goroutine to receive it
// Unbuffered channel is of zero capacity.

// The capacity of a buffered channel is the number of elements which that channel can hold. Capacity refers to the size of the buffer of the channel. The capacity of the channel can be specified during the creation of the channel while using the make function. The second argument is the capacity
// Capacity of unbuffered channel is always zero
// ch := make(chan int, 3)
// fmt.Printf("Capacity: %d\n", cap(ch))

// Builtin len() function can be used to get the length of a channel. The length of a channel is the number of elements that are already there in the channel. So length actually represents the number of elements queued in the buffer of the channel. Length of a channel is always less than or equal to the capacity of the channel.
// Length of unbuffered channel is always zero

// Close is an inbuilt function that can be used to close a channel. Closing of a channel means that no more data can we send to the channel.  Channel is generally closed when all the data has been sent and there's no more data to be send. Let's see a program
