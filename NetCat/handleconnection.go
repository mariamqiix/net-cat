package NetCat

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	firstTime := false
	name := ""
	Index := 0

	for {
		//at the beginning of each new connection, the user is asked to write his name. And print wlc message
		if !firstTime {
			conn.Write([]byte("Welcome to TCP-Chat!" + "\n"))
			for _, line := range WlcLogo {
				conn.Write([]byte(line + "\n"))
			}

			name = NameExistence(conn) //check if the name exists or not

			mutex.Lock()
			ClientsNames = append(ClientsNames, name) //append the client name
			Index = len(ClientsNames) - 1             //save the index of the client, so in case he wants to change his name
			mutex.Unlock()

			if name == "" { // we need to fix that
				conn.Write([]byte("Not a vaild name" + "\n")) //Control C problem --- user exit any way
				Exit(conn)
				return
			}

			if len(HistoryMessage) != 0 { //if the history message is not empty print it to the new client
				for _, message := range HistoryMessage {
					conn.Write([]byte(message + "\n"))
				}
			}
			//inform other clients, that new client join
			PrintMessage(conn, fmt.Sprint(colors[Index], name)+fmt.Sprint("\u001b[0m", " has joined our chat..."))

			firstTime = true
		}

		time.Sleep(time.Second)
		message, err := reader.ReadString('\n')

		if err != nil || message[:len(message)-1] == "exit" { //if user write exit, inform the thers and exit
			PrintMessage(conn, fmt.Sprint(colors[Index], name)+fmt.Sprint("\u001b[0m", " has left our chat..."))
			Exit(conn)
			return
		} else if message[:len(message)-1] == "--ChangeName" { //if user prifide this, change has name
			PrivesName := name // again make sure that the name is not exists
			name = NameExistence(conn)
			ClientsNames[Index] = name
			PrintMessage(conn, fmt.Sprint(colors[Index], PrivesName)+fmt.Sprint("\u001b[0m", " change has name to "+fmt.Sprint(colors[Index], name)))
		} else if len(message[:len(message)-1]) != 0 { //if message not empty send it through the chat
			currentTime := time.Now()

			messageB := fmt.Sprint("\u001b[0m", "["+currentTime.Format("2006-01-02 15:04:05")+"][") + fmt.Sprint(colors[Index], name) + fmt.Sprint("\u001b[0m", "]:"+message[:len(message)-1])
			mutex.Lock()
			HistoryMessage = append(HistoryMessage, messageB) //append the message to historymessages so if new client join we can show it to him
			mutex.Unlock()

			PrintMessage(conn, messageB)
		}
	}
}
