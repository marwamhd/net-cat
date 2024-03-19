package handlers

import (
	"fmt"
	"net"
	"sync"

	"netcat/Functions/natheerspretty"
)

var Connections []Connection

type Connection struct {
	Name string
	Conn net.Conn
}

var History []string

func HandleConnection(conn net.Conn) {
	done := make(chan bool, 1)
	welcome := "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n"
	var name string
	var err error
	go func() {
		SendMessageToWithChannel(conn, welcome, done)
	}()

	<-done

	name, err = GetClientName(conn)
	if err != nil {
		return
	}

	if len(Connections) >= 10 {
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
