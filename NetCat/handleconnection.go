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

	messageLimit := 10               // Maximum number of messages allowed per second
	perSecond := time.Second       // Time window duration (1 second)

	messageCounter := 0            // Counter to track the number of messages sent
	previousSecond := time.Now()   // Variable to store the previous second

	for {
		//at the beginning of each new connection, the user is asked to write his name. And print wlc message
		if !firstTime {
			conn.Write([]byte("Welcome to TCP-Chat!" + "\n"))
			for _, line := range WlcLogo {
				conn.Write([]byte(line + "\n"))
			}

			name = NameExistence(conn) //check if the name exists or not

			if name == "" {
				return
			}

			client := Client{conn: conn}

			mutex.Lock()

			clients = append(clients, client) //append the client to the client list so that he can receive new messages from the others
			ClientsNames = append(ClientsNames, name) //append the client name
			Index = len(ClientsNames) - 1             //save the index of the client, so in case he wants to change his name
			
			mutex.Unlock()

			if len(HistoryMessage) != 0 { //if the history message is not empty print it to the new client
				for _, message := range HistoryMessage {
					conn.Write([]byte(message + "\n"))
				}
			}
			//inform other clients, that new client join
			PrintMessage(conn, fmt.Sprint(colors[Index], name)+fmt.Sprint("\u001b[0m", " has joined our chat..."))

			firstTime = true
		}

		time.Sleep(3 * time.Millisecond)

		message, err := reader.ReadString('\n')

		currentSecond := time.Now()
		if currentSecond.Sub(previousSecond) >= perSecond {
			messageCounter = 0        // Reset the message counter
			previousSecond = currentSecond
		}

		// Check if the message limit is reached
		if messageCounter >= messageLimit {
			PrintMessage(conn, fmt.Sprint(colors[Index], name)+fmt.Sprint("\u001b[0m", " has left our chat..."))
			Exit(conn)
			return
			time.Sleep(time.Millisecond * 100)  // Sleep for a short duration before checking again
			continue
		}

		messageCounter++

		if err != nil || message[:len(message)-1] == "exit" || messageCounter > 10 { //if user write exit, inform the thers and exit
			PrintMessage(conn, fmt.Sprint(colors[Index], name)+fmt.Sprint("\u001b[0m", " has left our chat..."))
			Exit(conn)
			return
		} else if message[:len(message)-1] == "--ChangeName" { //if user prifide this, change has name
			PrivesName := name // again make sure that the name is not exists
			name = NameExistence(conn)
			ClientsNames[Index] = name
			PrintMessage(conn, fmt.Sprint(colors[Index], PrivesName)+fmt.Sprint("\u001b[0m", " change has name to "+fmt.Sprint(colors[Index], name)))
		} else if len(message[:len(message)-1]) != 0 && len(message[:len(message)-1]) < 500 && CheckLetters(message[:len(message)-1]) { //if message not empty send it through the chat
			currentTime := time.Now()

			messageB := fmt.Sprint("\u001b[0m", "["+currentTime.Format("2006-01-02 15:04:05")+"][") + fmt.Sprint(colors[Index], name) + fmt.Sprint("\u001b[0m", "]:"+message[:len(message)-1])
			mutex.Lock()
			HistoryMessage = append(HistoryMessage, messageB) //append the message to historymessages so if new client join we can show it to him
			mutex.Unlock()

			PrintMessage(conn, messageB)
		}
	}
}

func CheckLetters(s string) bool {
	for i:=0; i<len(s); i++ {
		if s[i] < 32 {
			return false
		}
	}
	return true
}
