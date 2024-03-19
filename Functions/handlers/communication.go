package handlers

import (
	"fmt"
	"net"
	"netcat/Functions/natheerspretty"
)

// function to broadcast message to selected client
func SendMessageTo(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

func handleClientMessage(client Connection, message []byte) {
	//print the message to the server

	//check if the message is empty
	if IsEmpty(message) {
		return
	}
	fmt.Println("the message is: ", string(message)+" and the length is: ", len(message))
	History = append(History, client.Name, ": ", string(message)+"\n")
	//broadcast the message to all clients
	BroadcastMessage(client.Name, ": "+string(message))
}

func BroadcastMessage(clientName, message string) {
	//broadcast the message to all clients
	for _, connection := range Connections {
		SendMessageTo(connection.Conn, clientName+message+"\n")
	}
}

func IsEmpty(message []byte) bool {
	if len(message) <= 0 {
		fmt.Println(natheerspretty.RGBify(255, 0, 0, "the message is empty"))
		return true
	}

	for i := 0; i < len(message); i++ {
		//check if the message is all spaces
		if message[i] != 32 {
			return false
		}
	}
	fmt.Println(natheerspretty.RGBify(255, 0, 0, "the message is full of spaces"))

	return true
}
