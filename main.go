package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
	// "math/rands" // to make username+random number
)

type Request struct {
	Client *User
	ChannelName string 
}

type User struct {
	Username string
	Nickname string
	Password string
	Message	chan Message
	Channel  ChatChannel
}

type ChatServer struct {
	AddUsr     chan User
	AddNick    chan User
	RemoveNick chan User
	NickMap    map[string]User
	Users      map[string]User
	Rooms      map[string]ChatChannel
	Create     chan ChatChannel
	Delete     chan ChatChannel
	UsrJoin    chan Request
	UsrLeave   chan Request
}

type ChatChannel struct {
	Name  string
	Users map[string]User

}

type Message struct {
	UserClient string // which could be Userclient User
	UserMessage string
}

func main() {
	ln, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	chatServer := &ChatServer{
		AddUsr:     make(chan User),
		AddNick:    make(chan User),
		RemoveNick: make(chan User),
		NickMap:    make(map[string]User),
		Users:      make(map[string]User),
		Rooms:      make(map[string]ChatChannel),
		Create:     make(chan ChatChannel),
		Delete:     make(chan ChatChannel),
		UsrJoin:    make(chan Request),
		UsrLeave:   make(chan Request),
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleconn(conn, chatServer)
	}
}

func handleconn(conn net.Conn, chatServer *ChatServer) {
	// err := conn.SetDeadline(time.Now().Add(20 * time.Second)) //Set up a timeout
	// if err != nil {
	// 	log.Println("CONNECTION TIMEOUT")
	// }
	fmt.Fprintf(conn, "NOTICE AUTH :*** Looking up your hostname...\nNOTICE AUTH :*** Found your hostname, welcome back\nNOTICE AUTH :*** Checking ident\nNOTICE AUTH :*** No identd (auth) response\n") //displays on the conn client
	var cl User
	io.WriteString(conn, "Enter your Username: ")
	scanner := bufio.NewScanner(conn) 
	scanner.Scan()
	Uname := scanner.Text()
	fmt.Println(Uname, "Just logged In")// need to store this in my database
	fmt.Fprintf(conn, "Your Username is %s\n", Uname)
	// scanner = bufio.NewScanner(conn) 
	// scanner.Scan()
	// Nname := scanner.Text() 
	// fmt.Println(Nname, "Is the Nickname to the Username: ", Uname)
	

	// cl.Username = Uname //
	if tmp, test := chatServer.Users[Uname]; test {
		cl = tmp
		io.WriteString(conn, "Enter your Password: ")
		scanner.Scan()
		pass := scanner.Text()
		for pass != cl.Password {
			io.WriteString(conn, "try again:\n")
			scanner.Scan()
			pass = scanner.Text()
		}

	} else {
		io.WriteString(conn, "Enter your Nickname\nPlease note you can change this at any time\nNickname: ")
		scanner.Scan()
		Nname := scanner.Text()
		for {
			if _, test := chatServer.NickMap[Nname]; test {
				io.WriteString(conn, "try again this Nickname is taken\n")
				scanner.Scan()
				Nname = scanner.Text()

			} else {
				break
			}

		}

		io.WriteString(conn, "Enter a Password for your account: ")
		scanner.Scan()
		pass := scanner.Text()
		tmp := User{
			Username:  Uname,
			Message: make(chan Message, 10),
			Nickname:   Nname,
			Password:     pass,
		}
		chatServer.AddUsr <- tmp
		cl = tmp
		fmt.Fprintf(conn, "You Are Welcome to Frozen IRC Server %s\n", Nname)
	}




	scanner1 := bufio.NewScanner(conn) //reads user input
	for scanner1.Scan() {              //loops through as long as it takes
		ln := scanner1.Text()                           //parses the input to ln
		fmt.Println(ln)                                //displays on the server
		fmt.Fprintf(conn, "I heard you say: %s\n", ln) //displays on the conn client
		
		

		//











	}
	defer conn.Close()

	fmt.Println(Uname, "Code Got To The Termination", "Or Exited")
}
