package main

import (
	"fmt"
	"net"

	webframework "github.com/kaestro/webframework/src"
)

func main() {
	engine := webframework.NewEngine()
	listener, err := engine.Start(webframework.DefaultAddress)
	if err != nil {
		fmt.Println("Error starting engine:", err)
		return
	}
	fmt.Println("Server listening on", webframework.DefaultAddress)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go func(c net.Conn) {
			engine.PrintRequest(c)
		}(conn)
	}
}
