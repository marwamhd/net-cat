package handlers

import (
	"fmt"
	"net"
	"time"

	"netcat/Functions/mainhelper"
)

// function to broadcast message to selected client
func SendMessageTo(conn net.Conn, message string) {
	conn.Write([]byte(message))
}

func SendMessageToWithChannel(conn net.Conn, message string, d chan bool) {
	conn.Write([]byte(message))
	d <- true
}

func handleClientMessage(client Connection, message []byte) {
	// check if the message is empty
	if mainhelper.IsEmpty(message) {
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := "[" + currentTime + "]:" + string(message)

	fmt.Println("[" + client.Name + "]" + formattedMessage)
	History = append(History, "["+client.Name+"]"+formattedMessage+"\n")
	// broadcast the message to all clients
	BroadcastMessageExceptSender(client.Name, formattedMessage)
	// SendMessageTo(client.Conn, "["+client.Name+"]"+"["+currentTime+"]:")
}

func BroadcastMessage(clientName, message string) {
	// broadcast the message to all clients
	for _, connection := range Connections {
		SendMessageTo(connection.Conn, "["+clientName+"]"+message+"\n")
	}
}

func BroadcastMessageExceptSender(clientName, message string) {
	// broadcast the message to all clients
	for _, connection := range Connections {
		if connection.Name != clientName {
			SendMessageTo(connection.Conn, "["+clientName+"]"+message+"\n")
		}
	}
}

func BroadcastMessageExceptSenderNewLine(clientName, message string) {
	// broadcast the message to all clients
	for _, connection := range Connections {
		if connection.Name != clientName {
			SendMessageTo(connection.Conn, "\n["+clientName+"]"+message+"\n")
		}
	}
}
