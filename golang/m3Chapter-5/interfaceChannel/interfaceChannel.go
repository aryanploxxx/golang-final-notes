package main

import (
	"fmt"
	"strconv"
)

// Messenger interface defines a method to get a message
type Messenger interface {
	Relay() string
}

// Message struct implements the Messenger interface
type Message struct {
	status string
}

// Relay returns the status of the Message
func (m Message) Relay() string {
	return m.status
}

// sendMessage sends a message to the channel with the given index
func sendMessage(messageChannel chan Messenger, index int) {
	message := &Message{
		status: "Task completed for index " + strconv.Itoa(index),
	}
	messageChannel <- message
}

func main() {
	// Create a channel for Messenger objects
	messageChannel := make(chan Messenger)

	// Launch 10 goroutines to send messages
	for i := 0; i < 10; i++ {
		go sendMessage(messageChannel, i)
	}

	// Wait for one message from the channel and print it
	select {
	case receivedMessage := <-messageChannel:
		fmt.Println(receivedMessage.Relay())
	}

	// Optional: receive another message to show a second example
	receivedMessage := <-messageChannel
	fmt.Println(receivedMessage.Relay())
}
