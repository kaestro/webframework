package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	serverAddress := "127.0.0.1:8080"

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	fmt.Printf("Connected to server %s\n", serverAddress)

	dataToSend := "print hello world"
	requestBytes := []byte(dataToSend)

	n, err := conn.Write(requestBytes)
	if err != nil {
		log.Fatalf("Failed to send data: %v", err)
	}
	fmt.Printf("Sent %d bytes to server\n", n)

	responseBuffer := make([]byte, 1024)
	nRead, err := conn.Read(responseBuffer)
	if err != nil && err != io.EOF {
		log.Fatalf("Failed to read response: %v", err)
	}
	fmt.Printf("Received %d bytes from server:\n%s\n", nRead, responseBuffer[:nRead])

	fmt.Println("Data sent and response received.")

	requestBytes = []byte("second request and close connection")
	n, err = conn.Write(requestBytes)
	if err != nil {
		log.Fatalf("Failed to send data: %v", err)
	}
	fmt.Printf("Sent %d bytes to server\n", n)

	responseBuffer = make([]byte, 1024)
	nRead, err = conn.Read(responseBuffer)
	if err != nil && err != io.EOF {
		log.Fatalf("Failed to read response: %v", err)
	}
	fmt.Printf("Received %d bytes from server:\n%s\n\n", nRead, responseBuffer[:nRead])

	conn.Close()

	conn, err = net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Connected to server %s\n", serverAddress)

	httpRequest := "GET / HTTP/1.1\r\nHost: 127.0.0.1:8080\r\nUser-Agent: my-go-client\r\nAccept: */*\r\n\r\n"
	requestBytes = []byte(httpRequest)

	n, err = conn.Write(requestBytes)
	if err != nil {
		log.Fatalf("Failed to send data: %v", err)
	}

	fmt.Printf("Sent %d bytes to server\n", n)

	responseBuffer = make([]byte, 1024)
	nRead, err = conn.Read(responseBuffer)
	if err != nil && err != io.EOF {
		log.Fatalf("Failed to read response: %v", err)
	}
	fmt.Printf("Received %d bytes from server:\n%s\n", nRead, responseBuffer[:nRead])

	fmt.Println("Data sent and response received.")
}
