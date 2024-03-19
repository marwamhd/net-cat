package initialconn

import (
	"fmt"
	"log"
	"net"
	"netcat/Functions/handlers"
	"netcat/Functions/mainhelper"
	"netcat/Functions/natheerspretty"
)

// This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

// type Connection struct {
// 	Name string
// 	Conn net.Conn
// }

func StartServer(ip string, port int) {
	listenAddress := net.JoinHostPort(ip, mainhelper.Itoa(port))
	tcpListen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatal(natheerspretty.RGBify(255, 0, 0, "error: ", err))
	}
	fmt.Println(natheerspretty.RGBify(0, 255, 0, "Server is listening on: ", listenAddress))

	for {
		conn, err := tcpListen.Accept()
		if err != nil {
			log.Fatal(natheerspretty.RGBify(255, 0, 0, "error: ", err))
		}

		// Handle the connection in a separate goroutine
		go handlers.HandleConnection(conn)

	}

	// Ensure tcpListen is closed when StartServer exits
	defer tcpListen.Close()
}
