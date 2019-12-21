package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	// "math/rands" // to make username+random number
)

type User struct {
	username string
	nickname string
	password string
	channel  ChatChannel
}

type ChatServer struct {
}

type ChatChannel struct {
	Name  string
	Users map[string]User
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
	scanner := bufio.NewScanner(conn) //reads user input
	for scanner.Scan() {              //loops through as long as it takes
		ln := scanner.Text()                           //parses the input to ln
		fmt.Println(ln)                                //displays on the server
		fmt.Fprintf(conn, "I heard you say: %s\n", ln) //displays on the conn client
	}
	defer conn.Close()

	fmt.Println("Code Got To The Termination")
}
