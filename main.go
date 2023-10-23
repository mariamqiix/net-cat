package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"net-cat/NetCat"
)


//Check if the we can use it from other PC -- should we use ip add insted of localhost? in handleconnection if name == ""
//test files for unit testing both the server connection and the client
//handle the errors from server side and client side
//project use channels or mutexes?
//Does the server produce logs about Clients activities?
//Are the server logs saved into a file?
//Is there more NetCat flags implemented?
//Does the project present a Terminal UI using JUST this package : https://github.com/jroimartin/gocui?

func main() {
	Port := "8989" //defualt port
 
	if len(os.Args) == 2 { //if user provide another port use it
		Port = os.Args[1]
	} else if len(os.Args) != 1 { //if arrgs > 1 amd not 2 , print error and exit
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	listener, err := net.Listen("tcp", "localhost:"+Port) //lisen the connection
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}

	defer listener.Close()

	fmt.Printf("Listening on the port :%s\n", Port)
	NetCat.EstabContact(listener)
}