package NetCat

import (
	"log"
	"net"
)

func EstabContact(listener net.Listener) {
	for { //establish and open the connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		client := Client{conn: conn}

		if len(clients) == 10 { //any new connection rejected if the existing connections reach 10
			conn.Write([]byte("Maximum number of clients reached. Connection rejected." + "\n"))
			log.Println("Maximum number of clients reached. Connection rejected.")
			conn.Close()
			continue
		}

		mutex.Lock()
		clients = append(clients, client) //append the client to the client list so that he can receive new messages from the others
		mutex.Unlock()

		go HandleConnection(conn) //go routines, to have synchronous connections
	}
}
