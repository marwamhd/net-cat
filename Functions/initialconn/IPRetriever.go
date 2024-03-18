package initialconn

import (
	"fmt"
	"net"
)

//This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

// make a function to get the ip address of the server
func GetIPAddress() (string, error) {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	localAddress := conn.LocalAddr().(*net.TCPAddr)
	conn.Close()
	return localAddress.IP.String(), nil
}
