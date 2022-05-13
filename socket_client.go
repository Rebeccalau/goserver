package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	fmt.Println("Started Client")

	request := `{"type": "testtype"}`

	_, _ = conn.Write([]byte(request))

	defer conn.Close()
}
