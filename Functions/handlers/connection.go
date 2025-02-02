package handlers

import (
	"fmt"
	"net"
	"sync"

	"netcat/Functions/mainhelper"
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
		fmt.Println(natheerspretty.RGBify(255, 0, 0, "Error: ", err))
		conn.Close()
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

	// send a welcome message to the client
	SendMessageTo(conn, "Welcome "+name+"!\n")

	if len(History) > 0 {
		for _, message := range History {
			SendMessageTo(conn, message)
		}
	}

	BroadcastMessageExceptSenderNewLine(client.Name, " has joined the chat")
	// currentTime := time.Now().Format("2006-01-02 15:04:05")
	// SendMessageTo(client.Conn, "["+client.Name+"]"+"["+currentTime+"]:")

	defer client.Conn.Close()
	fmt.Println(natheerspretty.RGBify(0, 255, 0, "Serving ", client.Conn.RemoteAddr().String()))
	for {
		buffer := make([]byte, 1024)
		n, err := client.Conn.Read(buffer)
		if err != nil || mainhelper.Signaltrapchecker(buffer[:n-1]) {
			handleClientDisconnect(client)
			return
		}

		handleClientMessage(client, buffer[:n-1])
	}
}
