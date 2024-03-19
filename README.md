# NetCat Project

This project aims to recreate the functionality of the NetCat command-line utility in a Server-Client Architecture. It allows for communication between a server and multiple clients over TCP or UDP connections.

## Features

- Server Mode: The server listens for incoming connections on a specified port.
- Client Mode: Clients can connect to a specified port and transmit information to the server.
- TCP and UDP Support: The project supports both TCP and UDP protocols for communication.
- Group Chat: The project enables group chat functionality, similar to the original NetCat.

## Usage

To use the NetCat project, follow these steps:

1. Clone the repository: `git clone https://github.com/yjawad/net-cat.git`
2. Navigate to the project directory: `cd net-cat`
3. build the server file: `go build -o netcatserver server.go`
4. run the server: `./netcatserver`
5. connect the client: `nc <ip> <port>`

