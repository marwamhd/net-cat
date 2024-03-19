package handlers

import (
	"fmt"
	"net"
	"netcat/Functions/natheerspretty"
	"time"
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

	currentTime := time.Now().Format("2015-07-27 12:44:45")
	formattedMessage := "[" + currentTime + "]:" + string(message)

	fmt.Println(formattedMessage)
	History = append(History, "["+client.Name+"]"+formattedMessage+"\n")
	//broadcast the message to all clients
	BroadcastMessage(client.Name, formattedMessage)
}

func BroadcastMessage(clientName, message string) {
	//broadcast the message to all clients
	for _, connection := range Connections {
		SendMessageTo(connection.Conn, "["+clientName+"]"+message+"\n")
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
