package handlers

import (
	"errors"
	"fmt"
	"net"
	"netcat/Functions/natheerspretty"
	"sync"
)

var Connections []Connection

type Connection struct {
	Name string
	Conn net.Conn
}

var History []string

func HandleConnection(conn net.Conn) {

	welcome := "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"

	go SendMessageTo(conn, welcome)
	name, err := GetClientName(conn)
	if err != nil {
		return
	}

	if len(Connections) >= 3 {
		SendMessageTo(conn, "error: the server is full\n")
		conn.Close()
		return
	}
	mutex := sync.Mutex{}

	mutex.Lock()
	client := Connection{Name: name, Conn: conn}

	Connections = append(Connections, client)
	mutex.Unlock()

	if len(History) > 0 {
		for _, message := range History {
			SendMessageTo(conn, message)
		}
	}

	//assume the first message is the name of the client
	//set the name of the client, -1 to remove the newline character

	BroadcastMessage(client.Name, " has joined the chat")

	defer client.Conn.Close()
	fmt.Println(natheerspretty.RGBify(0, 255, 0, "Serving ", client.Conn.RemoteAddr().String()))
	for {
		buffer := make([]byte, 1024)
		n, err := client.Conn.Read(buffer)
		if err != nil {
			handleClientDisconnect(client)
			return
		}

		handleClientMessage(client, buffer[:n-1])
	}

}

// function to set the name of the client
func GetClientName(conn net.Conn) (string, error) {
	name := ""
	var err error
	for IsEmpty([]byte(name)) {
		name, err = AwaitClientName(conn)
		if err != nil {
			return "", err
		}
		if IsEmpty([]byte(name)) {
			SendMessageTo(conn, "error: the name cannot be empty\n")
		}
		for _, connection := range Connections {
			if connection.Name == name {
				SendMessageTo(conn, "error: the name is already taken\n")
				name = ""
				break
			}
		}
	}

	return name, nil
}

// function to be awating for the client to send the name
func AwaitClientName(conn net.Conn) (string, error) {
	//fmt.Println("Enter your name!")
	SendMessageTo(conn, "[ENTER YOUR NAME]: ")
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(natheerspretty.RGBify(255, 0, 0, "guest has left without naming himself"))
		return "", errors.New("guest has left without naming himself")
	}
	return string(buffer[:n-1]), nil
}
