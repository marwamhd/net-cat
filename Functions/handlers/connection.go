package handlers

import (
	"fmt"
	"net"
	"netcat/Functions/natheerspretty"
)

var Connections []Connection

type Connection struct {
	Name string
	Conn net.Conn
}

func HandleConnection(client Connection) {
	//assume the first message is the name of the client
	//set the name of the client, -1 to remove the newline character

	actualClient := SetClientName(client)

	BroadcastMessage(actualClient.Name, " has joined the chat")

	defer actualClient.Conn.Close()
	fmt.Println(natheerspretty.RGBify(0, 255, 0, "Serving ", actualClient.Conn.RemoteAddr().String()))
	for {
		buffer := make([]byte, 1024)
		n, err := actualClient.Conn.Read(buffer)
		if err != nil {
			handleClientDisconnect(actualClient)
			return
		}

		handleClientMessage(actualClient, buffer[:n])
	}

}

// function to set the name of the client
func SetClientName(client Connection) Connection {
	name := ""
	for name == "" {
		name = AwaitClientName(client)
		if name == "" {
			SendMessageTo(client.Conn, "error: the name cannot be empty\n")
		}
		for _, connection := range Connections {
			if connection.Name == name {
				SendMessageTo(client.Conn, "error: the name is already taken\n")
				name = ""
				break
			}
		}
	}

	cid := 0
	//change name of the client in the connections
	for i, connection := range Connections {
		if connection.Conn == client.Conn {
			Connections[i].Name = name
			cid = i
			break
		}
	}
	return Connections[cid]
}

// function to be awating for the client to send the name
func AwaitClientName(client Connection) string {
	//fmt.Println("Enter your name!")
	SendMessageTo(client.Conn, "Enter your name!\n")
	buffer := make([]byte, 1024)
	n, err := client.Conn.Read(buffer)
	if err != nil {
		handleClientDisconnect(client)
		return ""
	}
	return string(buffer[:n-1])
}
