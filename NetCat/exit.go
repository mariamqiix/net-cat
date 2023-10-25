package NetCat

import "net"

func Exit(conn net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	// Remove the client from the list of clients and has name
	for i, c := range clients {
		if c.conn == conn {
			clients = append(clients[:i], clients[i+1:]...)
			if len(ClientsNames) > i {
				ClientsNames = append(ClientsNames[:i], ClientsNames[i+1:]...)
			}
			break
		}
	}

	return
}
