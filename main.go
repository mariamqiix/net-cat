package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

//To do list
//Control connections quantity >>> need to be check - try to have 10 connection and more and see what happen?
//same name issue!
// print this ugly :
//			_nnnn_
//         dGGGGMMb
//        @p~qp~~qMb
//        M|@||@) M|
//        @,----.JM|
//       JS^\__/  qKL
//      dZP        qKRb
//     dZP          qKKb
//    fZP            SMMb
//    HZM            MMMM
//    FqM            MMMM
//  __| ".        |\dS"qML
//  |    `.       | `' \Zq
// _)      \.___.,|     .'
// \____   )MMMMMP|   .'
//      `-'       `--'

type Client struct {
	conn net.Conn
	name string
}

var HistoryMessage []string

var clients []Client
var names []string

func main() {
	Server()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	firstTime := false
	name := ""
	reader := bufio.NewReader(conn)
	for {
		if !firstTime {
			conn.Write([]byte("Welcome to TCP-Chat!" + "\n"))
			name := NamesValdation(conn)

			if len(HistoryMessage) != 0 {
				for _, message := range HistoryMessage {
					conn.Write([]byte(message + "\n"))
				}
			}
			PrintMessage(conn, name[:len(name)-1]+" has joined our chat...")

			firstTime = true
		}

		message, err := reader.ReadString('\n')

		if err != nil {
			PrintMessage(conn, name[:len(name)-1]+" has left our chat...")
			// Remove the client from the list of clients
			for i, c := range clients {
				if c.conn == conn {
					clients = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			return
		}

		if len(message[:len(message)-1]) != 0 {
			currentTime := time.Now()

			messageB := "[" + currentTime.Format("2006-01-02 15:04:05") + "][" + name[:len(name)-1] + "]:[" + message[:len(message)-1] + "]"
			HistoryMessage = append(HistoryMessage, messageB)

			PrintMessage(conn, messageB)
		}
	}
}

func Server() {
	Port := "8989"

	//should we check if the user use port number less than 4 digit or use charcters or letters?
	if len(os.Args) == 2 {
		Port = os.Args[1]
	} else if len(os.Args) != 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	listener, err := net.Listen("tcp", "localhost:"+Port)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
	defer listener.Close()

	fmt.Printf("Listening on the port :%s\n", Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		client := Client{conn: conn}

		if len(clients) == 10 {
			log.Println("Maximum number of clients reached. Connection rejected.")
			conn.Close()
			continue
		}

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

func NamesValdation(conn net.Conn) string {
	conn.Write([]byte("[ENTER YOUR NAME]: " + "\n"))
	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n')
	for _, ClientName := range clients {
		if ClientName.name == name[:len(name)-1] {
			conn.Write([]byte("!!! THIS NAME ALREADY EXIST , CHOSE ANOTHER NAME !!!" + "\n"))
			name = NamesValdation(conn)
		}
	}
	return name[:len(name)-1]
}
