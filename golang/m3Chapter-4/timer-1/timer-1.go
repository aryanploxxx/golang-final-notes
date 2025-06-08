package main

import (
	"fmt"
	"sync"
	"time"
)

type TimeStruct struct {
	totalChanges int          // Keeps track of how many times the time was updated
	currentTime  time.Time    // Stores the current time
	rwLock       sync.RWMutex // Mutex for safe concurrent access
}

var TimeElement TimeStruct

func updateTime() {
	TimeElement.rwLock.Lock()            // Acquire write lock
	defer TimeElement.rwLock.Unlock()    // Release lock when function ends
	TimeElement.currentTime = time.Now() // Update the time
	TimeElement.totalChanges++           // Increment change counter
}

func main() {

	var wg sync.WaitGroup                // WaitGroup for goroutine synchronization
	TimeElement.totalChanges = 0         // Initialize counter
	TimeElement.currentTime = time.Now() // Set initial time

	timer := time.NewTicker(1 * time.Second) // Creates a ticker that "ticks" every 1 second
	//time.NewTicker(duration): Returns a new Ticker containing a channel that sends the current time at regular intervals
	// timer will return a struct which has a channel named C in it
	// NewTicker returns a new [Ticker] containing a channel that will send the current time on the channel after each tick. The period of the ticks is specified by the duration argument
	writeTimer := time.NewTicker(10 * time.Second) // Creates a ticker that "ticks" every 10 seconds
	endTimer := make(chan bool)                    // Channel to signal ending the goroutine

	wg.Add(1)

	go func() {
		for {
			select {
			case <-timer.C: // Every 1 second
				// timer.C: Returns the current time when the ticker ticks
				// Prints the current total changes and time
				fmt.Println(TimeElement.totalChanges, TimeElement.currentTime.String())

			case <-writeTimer.C: // Every 10 seconds
				updateTime() // Updates the time and increments counter

			case <-endTimer: // When endTimer receives a signal
				timer.Stop() // Stops the timer
				return       // Exits the goroutine
			}
		}
	}()

	wg.Wait()

	fmt.Println(TimeElement.currentTime.String())
}
