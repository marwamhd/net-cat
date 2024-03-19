package handlers

import (
	"errors"
	"fmt"
	"net"
	"netcat/Functions/mainhelper"
	"netcat/Functions/natheerspretty"
)

// function to set the name of the client
func GetClientName(conn net.Conn) (string, error) {
	name := ""
	var err error
	for mainhelper.IsEmpty([]byte(name)) {
		name, err = AwaitClientName(conn)
		if err != nil {
			return "", err
		}
		if mainhelper.IsEmpty([]byte(name)) {
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
