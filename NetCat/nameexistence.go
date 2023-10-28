package NetCat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func NameExistence(conn net.Conn) string { //this is recursive function, it will work until the user give valud name
	conn.Write([]byte("[ENTER YOUR NAME]: " + "\n"))
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n') //read the name from the user

	if err != nil {
		Exit(conn)
		return ""
	}

	if len(name) <= 1 {
		conn.Write([]byte(fmt.Sprint("\u001b[31m", "[!!! NAME CAN'T BE EMPTY, CHOOSE ANOTHER NAME !!!]: "+fmt.Sprint("\u001b[0m", "\n"))))
		return NameExistence(conn)
	} else if name[:len(name)-1] == "exit" || name[:len(name)-1] == "--ChangeName" {
		conn.Write([]byte(fmt.Sprint("\u001b[31m", "[!!! KEY WORD, CHOOSE ANOTHER NAME !!!]: "+fmt.Sprint("\u001b[0m", "\n"))))
		return NameExistence(conn)
	}

	if strings.TrimSpace(name[:len(name)-1]) == "" || !CheckLetters(name[:len(name)-1]) || len(strings.Join(strings.Fields(name[:len(name)-1]), " ")) > 20 {
		conn.Write([]byte(fmt.Sprint("\u001b[31m", "[!!! CHOOSE ANOTHER NAME !!!]: "+fmt.Sprint("\u001b[0m", "\n"))))
		return NameExistence(conn)
	}

	mutex.Lock()
	for _, Clientname := range ClientsNames { //check if it exists by looping in ClientName array
		if Clientname == name[:len(name)-1] {
			conn.Write([]byte(fmt.Sprint("\u001b[31m", "[!!! THIS NAME ALREADY EXISTS, CHOOSE ANOTHER NAME !!!]: "+fmt.Sprint("\u001b[0m", "\n"))))
			mutex.Unlock()
			return NameExistence(conn)
		}
	}

	mutex.Unlock()

	return strings.Join(strings.Fields(name[:len(name)-1]), " ") //if the loop end without re-calling the function return the name
}
