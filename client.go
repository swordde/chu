package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	fname := flag.String("filename", "file.aac", "path of the audio file")
	flag.Parse()

	file, err := os.Open(*fname)
	if err != nil {
		log.Fatal(err)
	}
	server()
}

func server() {
	// 1. Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	i := 0

	// 2. Take input from user
	for {
		//		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Say something: ")
		// o	text, _ := reader.ReadString('\n')

		// 3. Send to server
		conn.Write([]byte(fmt.Sprint(i)))
		i++

		// 4. Read server response
		reply := make([]byte, 1024)
		n, _ := conn.Read(reply)

		fmt.Println("Server replied:", string(reply[:n]))
	}
}
