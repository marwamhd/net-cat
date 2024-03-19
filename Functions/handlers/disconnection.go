package handlers

import (
	"fmt"
	"netcat/Functions/natheerspretty"
)

func handleClientDisconnect(client Connection) {
	fmt.Println("the length of the connections before was: ", len(Connections))
	defer client.Conn.Close()
	//remove the client from the connections
	removeClientFromConnections(client)
	fmt.Println("the length of the connections after is: ", len(Connections))
	//broadcast the message to all clients
	BroadcastMessage(client.Name, " has left the chat")
	//print the message to the server
	fmt.Println(natheerspretty.RGBify(255, 0, 0, client.Name, " has left the chat"))
}

package handlers

import (
	"fmt"
	"sync"
	"netcat/Functions/natheerspretty"
)

var disconnectionMutex sync.Mutex

func removeClientFromConnections(client Connection) {
	//remove the client from the connections global variable

	//add a mutex here
	disconnectionMutex.Lock()
	for i, connection := range Connections {
		if connection.Name == client.Name {
			Connections = append(Connections[:i], Connections[i+1:]...)
			break
		}
	}
	disconnectionMutex.Unlock()
}
