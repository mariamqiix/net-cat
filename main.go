package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Read command-line arguments
	args := os.Args[1:]
	port := "8989"

	if len(args) > 2 {
		fmt.Println("Usage: go run main.go [host] [port]")
		return
	} else if len(args) == 2 {
		port = args[1]
	}



	host := "ruqaya"

	// Connect to the specified host and port
	conn, err := net.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s:%s\n", host, port)

	// Start a goroutine to handle receiving messages from the server
	go receiveMessages(conn)

	// Read input from stdin and send it to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		// Send the input to the server
		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Println(err)
			break
		}
	}
}

// Function to receive messages from the server
func receiveMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}

		// Print the received message
		fmt.Print(strings.TrimSpace(message) + "\n")
	}
	// Close the connection when the server stops sending messages
	conn.Close()
}