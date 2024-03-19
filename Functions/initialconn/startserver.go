package initialconn

import (
	"fmt"
	"log"
	"net"
	"netcat/Functions/handlers"
	"netcat/Functions/mainhelper"
	"netcat/Functions/natheerspretty"
)

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
	//defer tcpListen.Close()
}
