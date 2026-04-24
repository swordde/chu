package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. Start server on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	for {
		// 2. Wait for client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 3. Handle client in new goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 4. Read data from client
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Client says:", message)

		// 5. Send response back
		response := "Hello from server"
		conn.Write([]byte(response))
	}
}
