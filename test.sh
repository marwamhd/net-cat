#!/bin/bash

# Run the Go server
go run server.go &

# Wait for the server to start
ping 10.30.18.80 -n 3 > nul

# Connect to the server with ncat in two new terminals
function connect_to_server {
     powershell.exe -Command "Start-Process cmd -ArgumentList \"/K ncat 10.30.18.80 8989\""
}

echo "Connecting to server..."
connect_to_server
sleep 2
connect_to_server