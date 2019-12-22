package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
	"strings"
	// "math/rands" // to make username+random number
)

// type Request struct {
// 	Client *User
// 	ChannelName string 
// }

type User struct {
	Username string
	Nickname string
	Password string
	// Channel  ChatChannel
}

	// var Users []User

// type ChatServer struct {
// 	AddUser 	chan	User
// 	AddNick 	chan	User
// 	RemoveNick chan 	User

// }

// type ChatChannel struct {
// 	Name  string
// 	Users map[string]User

// }

// type Message struct {
// 	UserClient string // which could be Userclient User
// 	UserMessage string
// }

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleconn(conn)
	}
}

func handleconn(conn net.Conn) {
	// err := conn.SetDeadline(time.Now().Add(20 * time.Second)) //Set up a timeout
	// if err != nil {
	// 	log.Println("CONNECTION TIMEOUT")
	// }
	fmt.Fprintf(conn, "NOTICE AUTH :*** Looking up your hostname...\nNOTICE AUTH :*** Found your hostname, welcome back\nNOTICE AUTH :*** Checking ident\nNOTICE AUTH :*** No identd (auth) response\n") //displays on the conn client
	io.WriteString(conn, "Enter your Username: ")
	scanner := bufio.NewScanner(conn) 
	scanner.Scan()
	Uname := scanner.Text()
	fmt.Println(Uname)// need to store this in my database
	fmt.Fprintf(conn, "Your Username is %s\n", Uname)
	io.WriteString(conn, "Enter your Nickname\nPlease note you can change this at any time\nNickname:")
	scanner = bufio.NewScanner(conn) 
	scanner.Scan()
	Nname := scanner.Text()
	fmt.Println(Nname)
	fmt.Fprintf(conn, "You Are Welcome to Frozen IRC Server %s\n", Nname)
	io.WriteString(conn, "Enter your Password: ")
	scanner.Scan()
	// Pass := scanner.Text()
	

	// scanner1 := bufio.NewScanner(conn) //reads user input ---------------before it enters loop
	// for scanner1.Scan() {              //loops through as long as it takes
	// 	ln := scanner1.Text()                           //parses the input to ln
	// 	fmt.Println(ln)                                //displays on the server
	// 	fmt.Fprintf(conn, "I heard you say: %s\n", ln) //displays on the conn client


			// instructions
		io.WriteString(conn, "\r\nBasic Input Instructions\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tPASS secretpasswordhere \r\n"+		//Sets the password of the user
		"\tNICK <nickname> \r\n"+				//sets the nickname of the user if not set and if set, to be changed
		"\tUSER <username> \r\n"+				//Sets the username
		"\tJOIN <#channel> \r\n"+				//Creates a channel if not exist, and if exists, join 
		"\tPART <#channel> \r\n"+				//Makes user leaves a channel
		"\tNAMES <#channel> \r\n"+				//lists all the Nicknames of users in a server
		"\tLIST <#channel> \r\n"+				//lists all channel or current status of channel
		"\tPRIVMSG <nickname>/<channel> \r\n"+	//sends a message to another user or channel 

		"\tGET fav \r\n\r\n\r\n")

	









		data := make(map[string]string)
		scanner1 := bufio.NewScanner(conn)
		for scanner1.Scan() {
			ln := scanner1.Text()
			fs := strings.Fields(ln)
			// logic
			if len(fs) < 1 {
				continue
			}
			switch fs[0] {
			case "GET":
				k := fs[1]
				v := data[k]
				fmt.Fprintf(conn, "%s\r\n", v)
			case "SET":
				if len(fs) != 3 {
					fmt.Fprintln(conn, "EXPECTED VALUE\r\n")
					continue
				}
				k := fs[1]
				v := fs[2]
				data[k] = v
			case "DEL":
				k := fs[1]
				delete(data, k)
			default:
				fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
				continue
			}









	}
	defer conn.Close()

	fmt.Println("Code Got To The Termination")
}
