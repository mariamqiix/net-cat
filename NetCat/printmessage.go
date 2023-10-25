package NetCat

import (
	"log"
	"net"
)

func PrintMessage(conn net.Conn, message string) { //this func to print the message to other users
	mutex.Lock()
	defer mutex.Unlock()

	for _, otherClient := range clients { //looping through client array, and send the message
		if otherClient.conn != conn { //if the client is not he who has send the message, send it to it, other ways ignore
			_, err := otherClient.conn.Write([]byte(message + "\n"))
			if err != nil {
				log.Println("Error sending message to client:", err)
			}
		}
	}
}
