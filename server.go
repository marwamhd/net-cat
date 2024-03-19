package main

import (
	"log"
	"netcat/Functions/initialconn"
	"netcat/Functions/mainhelper"
	"netcat/Functions/natheerspretty"
	"os"
)

func main() {
	args := os.Args[1:]
	port := 8989

	if len(args) > 1 {
		log.Fatal(natheerspretty.RGBify(255, 0, 0, "[USAGE]: ./TCPChat $port"))
	} else if len(args) == 1 {
		portNumber, err := mainhelper.Atoi(args[0])
		if err != nil {
			log.Fatal(natheerspretty.Redify("error: ", err))
		} else if portNumber < 0 || portNumber > 65535 {
			log.Fatal(natheerspretty.Redify("error: invalid port number: the port number must be between 0 and 65535"))
		}
		port = portNumber
	}

	ip, err := initialconn.GetIPAddress()
	if err != nil {
		natheerspretty.Redify(err)
		return
	}
	// fmt.Println("Server IP Address:", ip)
	// fmt.Println("Server is running on port:", mainhelper.Itoa(port))

	//now start listenting for tcp connection on the specified port
	initialconn.StartServer(ip, port)

}
