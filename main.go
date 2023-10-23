package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Client struct {
	conn net.Conn
	name string
}

var HistoryMessage []string
var clients []Client
var ClientsNames []string

func main() {
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
			conn.Write([]byte("Maximum number of clients reached. Connection rejected." + "\n"))
			log.Println("Maximum number of clients reached. Connection rejected.")
			conn.Close()
			continue
		}

		clients = append(clients, client)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	WlcLogo := []string{
		"          _nnnn_",
		"         dGGGGMMb",
		"        @p~qp~~qMb",
		"        M|@||@) M|",
		"        @,----.JM|",
		"       JS^\\__/  qKL",
		"      dZP        qKRb",
		"     dZP          qKKb",
		"    fZP            SMMb",
		"    HZM            MMMM",
		"    FqM            MMMM",
		"  __| \".        |\\dS\"qML",
		"  |    `.       | `' \\Zq",
		" _)      \\.___.,|     .'",
		" \\____   )MMMMMP|   .'",
		"      `-'       `--'",
	}

	reader := bufio.NewReader(conn)
	firstTime := false
	name := ""

	for {
		if !firstTime {
			conn.Write([]byte("Welcome to TCP-Chat!" + "\n"))
			for _, line := range WlcLogo {
				conn.Write([]byte(line + "\n"))
			}
			conn.Write([]byte("[ENTER YOUR NAME]: " + "\n"))

			name, _ = reader.ReadString('\n')

			for !NameExitene(name) {
				conn.Write([]byte("[!!! THIS NAME ALREADY EXIST, CHOSE ANOTHER NAME !!!]: " + "\n"))
				name, _ = reader.ReadString('\n')
			}

			ClientsNames = append(ClientsNames, name)

			if name == "" {
				Exit(conn)
				return
			}

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
			Exit(conn)
			return
		}

		if len(message[:len(message)-1]) != 0 {
			currentTime := time.Now()

			messageB := "[" + currentTime.Format("2006-01-02 15:04:05") + "][" + name[:len(name)-1] + "]:" + message[:len(message)-1]
			HistoryMessage = append(HistoryMessage, messageB)

			PrintMessage(conn, messageB)
		}
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

func NameExitene(name string) bool {
	for _, Clientname := range ClientsNames {
		if Clientname == name {
			return false
		}
	}
	return true
}

func Exit(conn net.Conn) {
	// Remove the client from the list of clients
	for i, c := range clients {
		if c.conn == conn {
			clients = append(clients[:i], clients[i+1:]...)
			ClientsNames = append(ClientsNames[:i], ClientsNames[i+1:]...)
			break
		}
	}
}
