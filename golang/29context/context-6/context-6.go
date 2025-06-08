package main

import "fmt"

func main() {
	tasks := []string{"task1", "task2", "task3"}

	ch := make(chan string, 3)

	// Sending tasks to the channel
	for _, task := range tasks {
		ch <- task
	}

	// Close the channel after sending all tasks
	close(ch)

	// Start the worker goroutine
	go worker(ch)

	// Allow the worker goroutine to complete
	fmt.Scanln()
}

// Worker function to process tasks
func worker(ch chan string) {
	for task := range ch {
		fmt.Println(task)
	}
}

/*

	How Closing a Channel Works:
		Readability of a Closed Channel:

		A closed channel can still be read by a goroutine until all buffered values are consumed.
		Once the buffer is empty, reading from a closed channel will return the zero value of the channel's type (e.g., "" for a chan string) and a second return value of false (if explicitly checked).
		range Behavior:

		When you use for t := range ch, the loop continues to read values from the channel until it is both empty and closed.
		Once all buffered values have been read and the channel is closed, the range loop terminates gracefully.

	Why Does the Worker Still Work?
		In your program:

		The main function populates the channel with tasks (e.g., "task1", "task2", "task3") before closing it.
		The worker goroutine starts after the tasks are already sent, and it uses for t := range ch to process the tasks.
		When the worker reads from the channel:
		It first consumes the buffered values.
		Once all buffered values are read and the channel is empty, the range loop detects that the channel is closed and exits gracefully.

*/
