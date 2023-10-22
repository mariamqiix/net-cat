package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

//To do list 
// Control connections quantity
// If a Client joins the chat, all the previous messages sent to the chat must be uploaded to the new Client.
// If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.
//If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.
//If a Client leaves the chat, the rest of the Clients must not disconnect.
//If there is no port specified, then set as default the port 8989. Otherwise, program must respond with usage message: [USAGE]: ./TCPChat $port

type Client struct {
	conn net.Conn
	name string
}

var HistoryMessage []string

var clients []Client

func main() {
	Server()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	firstTime := false
	name := ""
	for {
		if !firstTime {
			conn.Write([]byte("Welcome..." + "\n"))
			conn.Write([]byte("Enter a message to send (or 'quit' to exit): " + "\n"))
			conn.Write([]byte("Write your name: " + "\n"))

			name, _ = reader.ReadString('\n')
			firstTime = true
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading:", err)
			return
		}

		if len(message[:len(message)-1]) != 0 {
			currentTime := time.Now()
			timeString := currentTime.Format("2006-01-02 15:04:05")

			messageB := "[" + timeString + "][" + name[:len(name)-1] + "]:[" + message[:len(message)-1] + "]"
			HistoryMessage = append(HistoryMessage, messageB)

			fmt.Println(messageB)

			if message == "quit" { // fix that
				break
			}

			PrintMessage(conn, messageB)
		}
	}
}

func Server(){
	listener, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
	defer listener.Close()

	fmt.Println("Server started. Waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		client := Client{conn: conn}
		clients = append(clients, client)
		go handleConnection(conn)
	}
}

func PrintMessage(conn net.Conn, message string) {
	for _, otherClient := range clients {
		if otherClient.conn != conn {
			_, err := otherClient.conn.Write([]byte(message + "\n"))
			if err != nil {
				log.Println("Error sending message to client:", err)
			}
		}
	}
}
