package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	config2 "start_server/config"
)

type SocketRequest struct {
	RequestType string `json:"type"`
}

func main() {
	arguments := os.Args[1]
	config := config2.NewConfiguration(arguments)
	ln, _ := net.Listen("tcp", "localhost:8000")

	defer ln.Close()
	for {
		conn, _ := ln.Accept()

		message, _ := bufio.NewReader(conn).ReadString('\n')
		go processMessage(config, message)
	}
}

func processMessage(config config2.Configuration, message string) {
	var parsedRequest SocketRequest

	json.Unmarshal([]byte(message), &parsedRequest)

	switch parsedRequest.RequestType {
	case "testtype":
		{
			fmt.Println("Testing TestType")
			fmt.Println("Config got here", config.Environment)
		}
	default:
		{
			fmt.Println("Default")
		}
	}
}
