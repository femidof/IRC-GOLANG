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
	Channel  ChatChannel
}

type ChatServer struct {
	AddUser 	chan	User
	AddNick 	chan	User
	RemoveNick chan 	User

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




	scanner1 := bufio.NewScanner(conn) //reads user input
	for scanner1.Scan() {              //loops through as long as it takes
		ln := scanner1.Text()                           //parses the input to ln
		fmt.Println(ln)                                //displays on the server
		fmt.Fprintf(conn, "I heard you say: %s\n", ln) //displays on the conn client
	}
	defer conn.Close()

	fmt.Println("Code Got To The Termination")
}
