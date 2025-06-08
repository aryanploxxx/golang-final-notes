/*
	Transmit sigmals between goroutines using cond
	Like a waiting room, where a goroutine sleeps/waits till it gets a signal
	Condition is just a signal that something has happend
	Condition does not provide any information or data

	Cond implements a condition varaible, a point for goroutines waiting for or announcing the occurence of an event
	It is a signal that announces the occurance of a condition being met

	Naive approach is to use infinite loop, with checking of the condition being met after every iteration
	-> This is not an efficient way, as it will keep on checking the condition, even if it is not met
	Better way is making a 'waiting goroutine' which does nothing until it gets the signal

	Locks must be help when chaging the condition or calling the method

	Locker - used to work on the condition
	Wait()- cause the calling routine to block until another goroutine signals a condition, that the condition has been met
	Signal()- wakes up one goroutine that is waiting on the condition
		-> signal can be implemented with channles but implmenting broadcast is not easy
	Broadcast()- wakes up all goroutines that are waiting on the condition
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var integers = make([]int, 0, 10)

func main() {
	go add()
	// Starts the add function in a separate goroutine.
	time.Sleep(5 * time.Second) // fake delay
}

func add() {
	for i := 0; i < 5; i++ {
		cond.L.Lock()
		for len(integers) == 2 {
			cond.Wait()
			fmt.Println("Len befire add: ", len(integers))
		}
		integers = append(integers, i)
		fmt.Println("Len after add: ", len(integers))
		go remove(1 * time.Second)
		cond.L.Unlock()
	}
}

func remove(delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Len before remove: ", len(integers))
	integers = integers[1:]
	fmt.Println("Len after remove: ", len(integers))
	cond.Signal()
}

/*

add() starts in a goroutine and begins adding integers to the integers slice.
If the slice length reaches 2, the add goroutine waits until the condition variable is signaled.
For each addition, the remove function is started in a separate goroutine, which removes an element after a 1-second delay.
The remove function signals the add function to resume if it was waiting.
The main function sleeps for 5 seconds, allowing the other goroutines to execute.
This demonstrates a producer-consumer pattern where add (producer) adds elements and remove (consumer) removes elements with synchronization.

*/
