package handlers

import (
	"fmt"
	"log"
	"net"
	"netcat/Functions/natheerspretty"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println(natheerspretty.RGBify(0, 255, 0, "Serving ", conn.RemoteAddr().String()))
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(natheerspretty.RGBify(255, 0, 0, "error: "), err)
			return
		}
		fmt.Print(string(buffer[:n]))
	}
}
