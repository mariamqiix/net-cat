package main

import (
	"fmt"
	"log"
	"net"
	"net-cat/NetCat"
	"os"
	"strconv"
)

//test files for unit testing both the server connection and the client

func main() {
	Port := "8989" //defualt port

	if len(os.Args) == 2 { //if user provide another port use it
		Port = os.Args[1]
		_, err := strconv.Atoi(Port)
		if err != nil || len(Port) != 4 || Port[0] == '-' {
			fmt.Println("[USAGE]: ./TCPChat $port")
			return
		}
	} else if len(os.Args) != 1 { //if arrgs > 1 and not 2 , print error and exit
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	fmt.Println(NetCat.GetIpAdd())

	listener, err := net.Listen("tcp", NetCat.GetIpAdd()+":"+Port) //lisen the connection
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}

	defer listener.Close()

	fmt.Printf("Listening on the port :%s\n", Port)
	NetCat.EstabContact(listener)
}
