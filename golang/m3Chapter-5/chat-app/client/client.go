package main

import (
	"bufio"   // Provides buffered I/O operations for reading and writing
	"fmt"     // Used for formatted I/O operations
	"net"     // Provides networking functionality like TCP connections
	"os"      // Used for command-line arguments and standard input/output
	"strings" // Provides string manipulation utilities
)

// Message struct (unused in this implementation) represents a user message
// with the message content and the username.
type Message struct {
	message string // Message content
	user    string // Username
}

// recvBuffer (unused in this implementation) is a buffer for receiving data.
var recvBuffer [140]byte

// listen continuously reads messages from the server and prints them.
func listen(conn net.Conn) {
	for {
		// Create a buffer to hold incoming messages from the server.
		messBuff := make([]byte, 1024)

		// Read data into the buffer from the connection.
		n, err := conn.Read(messBuff)
		if err != nil {
			// Handle connection errors (e.g., server closed connection).
			return
		}

		// Convert the received bytes to a string.
		message := string(messBuff[:n])

		// Trim any extra whitespace from the message.
		message = strings.TrimSpace(message)

		// Print the received message.
		fmt.Println(message)
		fmt.Print("> ") // Display prompt for user input.
	}
}

// talk reads user input from the terminal and sends it to the server.
func talk(conn net.Conn, mS chan Message) {
	for {
		// Create a buffered reader to read input from the terminal.
		command := bufio.NewReader(os.Stdin)

		// Display the prompt.
		fmt.Print("> ")

		// Read a line of input from the user.
		line, err := command.ReadString('\n')
		if err != nil {
			// If there's an error reading input, close the connection and exit.
			conn.Close()
			break
		}

		// Trim newline and extra whitespace from the input.
		line = strings.TrimRight(line, " \t\r\n")

		// Write the input to the server over the connection.
		_, err = conn.Write([]byte(line))
		if err != nil {
			// If there's an error sending the message, close the connection and exit.
			conn.Close()
			break
		}

		// Call a function that does nothing (placeholder for additional logic).
		doNothing(command)
	}
}

// doNothing is a placeholder function that currently does nothing.
func doNothing(bf *bufio.Reader) {
}

// main is the entry point of the program.
func main() {
	// Create a channel for messages (unused in this implementation).
	messageServer := make(chan Message)

	// Get the username from the command-line arguments.
	userName := os.Args[0]

	// Print a message indicating the user is connecting to the server.
	fmt.Println("Connecting to host as", userName)

	// Create a channel to detect when the client should close.
	clientClosed := make(chan bool)

	// Establish a TCP connection to the server.
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		// If the connection fails, print an error message and exit.
		fmt.Println("Could not connect to server!")
		return
	}

	// Send the username to the server.
	conn.Write([]byte(userName))

	// Create a buffer to read the server's introductory message.
	introBuff := make([]byte, 1024)

	// Read the introductory message from the server.
	n, err := conn.Read(introBuff)
	if err != nil {
		// If there's an error reading the intro message, exit.
		return
	}

	// Convert the received bytes into a string and print it.
	message := string(introBuff[:n])
	fmt.Println(message)

	// Start a goroutine to handle sending user messages.
	go talk(conn, messageServer)

	// Start a goroutine to handle receiving messages from the server.
	go listen(conn)

	// Block the main function until the client is closed.
	<-clientClosed
}
