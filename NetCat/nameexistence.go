package NetCat

import (
	"bufio"
	"fmt"
	"net"
)

func NameExistence(conn net.Conn) string { //this is recursive function, it will work until the user give valud name
	conn.Write([]byte("[ENTER YOUR NAME]: " + "\n"))
	reader := bufio.NewReader(conn)
	name, _ := reader.ReadString('\n') //read the name from the user
	if name != "" {
		name = name[:len(name)-1]
	}
	mutex.Lock()
	for _, Clientname := range ClientsNames { //check if it exists by looping in ClientName array
		if Clientname == name {
			conn.Write([]byte(fmt.Sprint("\u001b[31m", "[!!! THIS NAME ALREADY EXISTS, CHOOSE ANOTHER NAME !!!]: "+fmt.Sprint("\u001b[0m", "\n"))))
			mutex.Unlock()
			return NameExistence(conn)
		} else if name == "exit" || name == "--ChangeName" {
			conn.Write([]byte(fmt.Sprint("\u001b[31m", "[!!! KEY WORD, CHOOSE ANOTHER NAME !!!]: "+fmt.Sprint("\u001b[0m", "\n"))))
			mutex.Unlock()
			return NameExistence(conn)
		} else if name == "" {
			conn.Write([]byte(fmt.Sprint("\u001b[31m", "[!!! NAME CAN'T BE EMPTY, CHOOSE ANOTHER NAME !!!]: "+fmt.Sprint("\u001b[0m", "\n"))))
			mutex.Unlock()
			return NameExistence(conn)
		}
	}
	mutex.Unlock()
	return name //if the loop end without re-calling the function return the name
}
