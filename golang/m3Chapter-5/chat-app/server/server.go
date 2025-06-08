package main

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var connectionCount int
var messagePool chan (string)

const (
	INPUT_BUFFER_LENGTH = 140
)

type User struct {
	Name      string      // Username
	ID        int         // Unique ID (optional in this example)
	Initiated bool        // Indicates if the user is fully connected
	UChannel  chan []byte // Channel for receiving messages
	// UChannel: Each user has a channel for receiving messages. This ensures each user can process messages independently.
	Connection *net.Conn // TCP connection object
	// Connection: Stores the TCP connection for sending and receiving data.
}

func (u *User) Listen() {
	fmt.Println("Listening for", u.Name)
	for {
		select {
		case msg := <-u.UChannel:
			fmt.Println("Sending new message to", u.Name)
			fmt.Fprintln(*u.Connection, string(msg))

		}
	}
}

/*
	The ConnectionManager manages:
	- Server metadata (e.g., name).
	- All user connections.
	The Initiate() function initializes this struct when the server starts.
*/

type ConnectionManager struct {
	name      string
	initiated bool
}

func Initiate() *ConnectionManager {
	cM := &ConnectionManager{
		name:      "Chat Server 1.0",
		initiated: false,
	}

	return cM
}

// The evalMessageRecipient function checks if a message is intended for a specific user:
func evalMessageRecipient(msg []byte, uName string) bool {
	eval := true
	expression := "@"
	re, err := regexp.MatchString(expression, string(msg))
	if err != nil {

	}
	if re {
		eval = false
		pmExpression := "@" + uName
		pmRe, pmErr := regexp.MatchString(pmExpression, string(msg))
		if pmErr != nil {

		}
		if pmRe {
			eval = true
		}
	}
	return eval // True if the message is for the user
}

// The server continuously listens for new connections in the Listen method
// Each user has their own Listen method:
func (cM *ConnectionManager) Listen(listener net.Listener) {
	fmt.Println(cM.name, "Started")
	for {
		conn, err := listener.Accept() // Accepts a new connection
		if err != nil {
			fmt.Println("Connection error", err)
		}

		connectionCount++
		fmt.Println(conn.RemoteAddr(), "connected")

		user := User{Name: "anonymous", ID: 0, Initiated: false}
		// User Initialization: New users are initially assigned the username "anonymous" and later updated when they send their first message.
		Users = append(Users, &user) // Add user to the global list

		for _, u := range Users {
			fmt.Println("User online", u.Name)
		}
		fmt.Println(connectionCount, "connections active")

		go cM.messageReady(conn, &user) // Handle user messages concurrently
	}
}

// The messageReady function manages user interactions:
func (cM *ConnectionManager) messageReady(conn net.Conn, user *User) {
	uChan := make(chan []byte)

	for {
		buf := make([]byte, INPUT_BUFFER_LENGTH) // Buffer for incoming messages
		n, err := conn.Read(buf)                 // Read message from client
		if err != nil {                          // Handle disconnections
			conn.Close()
		}
		if n == 0 {
			conn.Close()
		}

		fmt.Println(n, "character message from user", user.Name)

		if !user.Initiated {
			fmt.Println("New User is", string(buf))
			user.Initiated = true
			user.UChannel = uChan
			user.Name = string(buf[:n]) // Set username
			user.Connection = &conn
			go user.Listen() // Start listening for messages

			minusYouCount := strconv.FormatInt(int64(connectionCount-1), 10)
			conn.Write([]byte("Welcome to the chat, " + user.Name + ", there are " + minusYouCount + " other users"))

		} else {
			sendMessage := []byte(user.Name + ": " + strings.TrimRight(string(buf), " \t\r\n"))
			for _, u := range Users {
				if evalMessageRecipient(sendMessage, u.Name) { // Broadcast or direct message
					u.UChannel <- sendMessage
				}
			}
		}

	}
}

var Users []*User

func main() {
	connectionCount = 0
	serverClosed := make(chan bool)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Could not start server!", err)
	}

	connManage := Initiate()
	go connManage.Listen(listener)

	<-serverClosed
}
