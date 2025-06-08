package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	go sum(ch)
	time.Sleep(time.Second * 1)
}

func sum(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}

/*
	In the above program,  we created a channel.   In the main function the send three values to the channel and after that, we closed the channel. Then we called the sum function and we passed the channel to that function. In the sum function, we did a for range loop over the channel.    After iterating over all the values in the channel the for range loop will exit  since the channel is closed
	Now the question which comes to the mind is that what happens if you don't close a channel in the main function.  Try commenting the line in which they are closing the channel. Now run the program.  It will also  output  deadlock because  for range loop will never finish in the sum function
*/
