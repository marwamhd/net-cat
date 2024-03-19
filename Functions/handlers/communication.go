package handlers

import (
	"fmt"
	"net"
	"netcat/Functions/mainhelper"
	"time"
)

// function to broadcast message to selected client
func SendMessageTo(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

func handleClientMessage(client Connection, message []byte) {

	//check if the message is empty
	if mainhelper.IsEmpty(message) {
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
