package webframework

import (
	"fmt"
	"io"
	"net"
)

type Engine interface {
	Start(address string) (net.Listener, error)
	Stop() error
	PrintRequest(conn net.Conn)
}

type webframework struct {
	listener net.Listener
}

func NewEngine() Engine {
	return &webframework{}
}

const (
	DefaultAddress = "127.0.0.1:8080"
)

func (e *webframework) Start(address string) (net.Listener, error) {
	if address == "" {
		address = DefaultAddress
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	e.listener = listener

	return e.listener, nil
}

func (e *webframework) Stop() error {
	return e.listener.Close()
}

func (e *webframework) PrintRequest(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Connection received")
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading from connection:", err)
			} else {
				fmt.Println("Client connection closed")
			}
			return
		}

		if n == 0 {
			continue
		}

		fmt.Printf("Received %d bytes: %s\n", n, string(buffer[:n]))

		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
		fmt.Printf("Sent back %d bytes\n", n)
	}
}
