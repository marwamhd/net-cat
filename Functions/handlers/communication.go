package handlers

import (
	"fmt"
	"net"
)

// function to broadcast message to selected client
func SendMessageTo(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

func handleClientMessage(client Connection, message []byte) {
	//print the message to the server
	fmt.Print(string(message))
	//broadcast the message to all clients
	BroadcastMessage(client.Name, ": "+string(message))
}

func BroadcastMessage(clientName, message string) {
	//broadcast the message to all clients
	for _, connection := range Connections {
		SendMessageTo(connection.Conn, clientName+message+"\n")
	}
}
