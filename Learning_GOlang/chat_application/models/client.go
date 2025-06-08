package models

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	clientsocket *websocket.Conn
	room         *Room
	sendchan     chan []byte
}

//in the read method
//take an instance that there are 2 users alice and bob
//alice types a message in her browser "hello bob" and sent it via http and it was sent using specialheader which switched the protocal to websocket so basically she sent via websocker
//now i need to forward this message to forwardchan of room in order for the room to broadcast it to all clients
//so we first need to readthe message from the socket and then send to forward channel

func (c *Client) Read() {

	for {
		_, msg, err := c.clientsocket.ReadMessage()
		if err != nil {
			fmt.Println("failed to retieve message")
			break
		}

		c.room.forward <- msg

	}
	c.clientsocket.Close()

}
func (c *Client) Write() {
	defer c.clientsocket.Close()

	for msg := range c.sendchan {
		if err := c.clientsocket.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}
