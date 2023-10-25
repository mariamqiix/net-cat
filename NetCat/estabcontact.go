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

		if len(clients) == 10 { //any new connection rejected if the existing connections reach 10
			conn.Write([]byte("Maximum number of clients reached. Connection rejected." + "\n"))
			log.Println("Maximum number of clients reached. Connection rejected.")
			conn.Close()
			continue
		}

		go HandleConnection(conn) //go routines, to have synchronous connections
	}
}
