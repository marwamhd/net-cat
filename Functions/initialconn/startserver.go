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

func StartServer(ip string, port int) {
	listenAddr := net.JoinHostPort(ip, mainhelper.Itoa(port))
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(natheerspretty.RGBify(255, 0, 0, "error: ", err))
	}
	fmt.Println(natheerspretty.RGBify(0, 255, 0, "Server is listening on: ", listenAddr))
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(natheerspretty.RGBify(255, 0, 0, "error: ", err))
		}

		// Handle the connection in a separate goroutine
		go handlers.HandleConnection(conn)
	}
}
