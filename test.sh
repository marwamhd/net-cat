#!/bin/bash

# Run the Go server
go run server.go &

# Wait for the server to start
ping ip -n 3 > nul

# Connect to the server with ncat in two new terminals
function connect_to_server {
     powershell.exe -Command "Start-Process cmd -ArgumentList \"/K ncat ip 8989\""
}

echo "Connecting to server..."
connect_to_server
sleep 2
connect_to_server