package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"io"
	"strings"
	"time"
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
	Status 	int 
	// Channel  ChatChannel
}

	// var Users []User
type ChatChannel struct {
	Name string
	Description string
	Users map[string]User
}

type ChatServer struct {
	// Users []User
	Channels map[string]ChatChannel
}

type ChatUsers struct {
	Name  string
	Users []User

}

// type Message struct {
// 	UserClient string // which could be Userclient User
// 	UserMessage string
// }

func main() {
	usr1 := User{"Femi", "Fem", "0000", 0}
	usr2 := User{"Victoria", "Ria", "1234", 0}
	tinder := ChatUsers{Name: "Tinder"}

	tinder.Users = append(tinder.Users, usr1)
	tinder.Users = append(tinder.Users, usr2)
	fmt.Println(tinder.Users)
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
		go handleconn(conn, tinder)
	}
}

func handleconn(conn net.Conn, tinder ChatUsers) {
	// err := conn.SetDeadline(time.Now().Add(20 * time.Second)) //Set up a timeout
	// if err != nil {
	// 	log.Println("CONNECTION TIMEOUT")
	// }
	// ch1 := ChatChannel{Name: "Welcome", Description: "Welcoming you to Frozen Server", Users: ""}
	// var chatserver ChatServer
	// chatserver.Channels.Name := channels{ Name: West ,Description: Westwingers , Users: "Femi","Victoria"}
	fmt.Fprintf(conn, "NOTICE AUTH :*** Looking up your hostname...\nNOTICE AUTH :*** Found your hostname, welcome back\nNOTICE AUTH :*** Checking ident\nNOTICE AUTH :*** No identd (auth) response\n") //displays on the conn client
	
	var Uname , Nname, pass string
	fmt.Println(Uname)// need to store this in my database
	
	// use new to create channel struct pointer!
	ch := new(ChatChannel)
	ch.Name = "room1"
	fmt.Println(ch)

	serverNew := new(ChatServer)
	serverNew.Channels = make(map[string]ChatChannel)
	serverNew.Channels["r1"] = *ch
	
	for key, value := range serverNew.Channels {
		fmt.Println("chatroom name = ", key, "channel struct", value)
	}






	
	// Pass := scanner.Text()
	

	// scanner1 := bufio.NewScanner(conn) //reads user input ---------------before it enters loop
	// for scanner1.Scan() {              //loops through as long as it takes
	// 	ln := scanner1.Text()                           //parses the input to ln
	// 	fmt.Println(ln)                                //displays on the server
	// 	fmt.Fprintf(conn, "I heard you say: %s\n", ln) //displays on the conn client
	auth := 0

here:
			// instructions
		io.WriteString(conn, "\r\nBasic Input Instructions\r\n\r\n"+
		"USE:\r\n"+
		"\tPASS <NICK> <USER> \r\n"+		//Sets the password of the user
		"\tNICK <nickname> \r\n"+				//sets the nickname of the user if not set and if set, to be changed
		"\tUSER <username> \r\n"+				//Sets the username
		"\tJOIN <#channel> \r\n"+				//Creates a channel if not exist, and if exists, join 
		"\tPART <#channel> \r\n"+				//Makes user leaves a channel
		"\tNAMES <#channel> \r\n"+				//lists all the Nicknames of users in a server
		"\tLIST <#channel> \r\n"+				//lists all channel or current status of channel
		"\tPRIVMSG <nickname>/<channel> \r\n\r\n\r\n\r\n")	//sends a message to another user or channel 

	





		// master := 0


		// data := make(map[string]string)
		scanner1 := bufio.NewScanner(conn)
		for scanner1.Scan() {
			ln := scanner1.Text()
			fs := strings.Fields(ln)
			// logic
			if len(fs) < 1 {
				continue
			}
			switch fs[0] {
				//This is to watch how to work with my setting values
			// case "GET":
			// 	k := fs[1]
			// 	v := data[k]
			// 	fmt.Fprintf(conn, "%s\r\n", v)
			// case "SET":
			// 	if len(fs) != 3 {
			// 		fmt.Fprintln(conn, "EXPECTED VALUE\r")
			// 		continue
			// 	}
			// 	k := fs[1]
			// 	v := fs[2]
			// 	data[k] = v
			// case "DEL":
			// 	k := fs[1]
			// 	delete(data, k)
			// case  "USER":
			// 	if len(fs) != 2 {
			// 		fmt.Fprintln(conn, "Ambiguous Value")
			// 		continue
			// 	}
			// 	if master == 0 { 	//To know if User name has been set before
			// 		io.WriteString(conn, "Enter your new Password: ")	//new User
			// 		scanner.Scan()
			// 		pass := scanner.Text()
			// 		v := fs[1]
			// 		Uname = v
			// 		new_user := User{Username: v, Password: pass, Nickname: v, Status:1}
			// 		tinder.Users = append(tinder.Users, new_user)
			// 		master += 1
			// 	} else {
			// 		// len(tinder.Users[])
			// 			for i:=0; i < len(tinder.Users); i++  {	// Old User Confirm your pass word
			// 				if	tinder.Users[i].Username == fs[1] {
			// 				io.WriteString(conn, "Enter your Old Password: ")	//new User
			// 				scanner.Scan()
			// 				pass := scanner.Text()
			// 				if	tinder.Users[i].Password == pass{
								
			// 					fmt.Fprintln(conn,"Welcome Back")
			// 					Uname = tinder.Users[i].Username
			// 				} else {
										
			// 					fmt.Fprintln(conn,"User still online... Not your account..")
			// 					conn.Close()
			// 				}

			// 		 } else {
			// 			fmt.Fprintln(conn,"User still online... Not your account..")
			// 			conn.Close()
			// 		}
			// 	}


			// 	}
				
				
			case  "PASS":
				if len(fs) != 3 && fs[2] != "USER" && fs[1] != "NICK" {
					fmt.Fprintln(conn, "Ambiguous Value")
					break
				}
				io.WriteString(conn, "Enter your Username: ")
				scanner := bufio.NewScanner(conn) 
				scanner.Scan()
				Uname = scanner.Text()	// Username
				
				io.WriteString(conn, "Enter your Nickname: ")
				scanner = bufio.NewScanner(conn) 
				scanner.Scan()
				Nname = scanner.Text()

				io.WriteString(conn, "Enter your Password: ")
				scanner = bufio.NewScanner(conn) 
				scanner.Scan()
				pass = scanner.Text()

				for i:=0; i < len(tinder.Users); i++  {
					if	tinder.Users[i].Username == Uname {
						fmt.Fprintln(conn,"Username Exists, Checking if Password Matches")
						if tinder.Users[i].Password == pass {
							fmt.Fprintln(conn,"Welcome Back")
							tinder.Users[i].Status = 1
							tinder.Users[i].Nickname = Nname
							auth+=1
							time.Sleep(10 * time.Second)

							goto here;
						} else {
							fmt.Fprintln(conn,"Password Incorrect")
							auth =0
							time.Sleep(10 * time.Second)
							goto here;
						}
						
					}				
				}
				fmt.Fprintln(conn,"Creating Username: ", Uname)
				new_user := User{Username: Uname, Password: pass, Nickname: Nname, Status:1}
				tinder.Users = append(tinder.Users, new_user)
				auth+=1

			case  "NICK":
				if auth == 0 {
					fmt.Fprintln(conn,"You are currently not logged In")
					time.Sleep(10 * time.Second)
					goto here;
				}
				if len(fs) == 2 {
					for i:=0; i < len(tinder.Users); i++ {
						if tinder.Users[i].Username == Uname {
							tinder.Users[i].Nickname = fs[1]
							Nname = fs[1]
							fmt.Fprintln(conn, "Nickname Changed To ", Nname)
							break
						}
					}
				} else {
					fmt.Fprintln(conn, "Ambiguous Value")
					fmt.Fprintln(conn, "Nickname not changed")
			}
			case  "JOIN":
				if auth == 0 {
					time.Sleep(10 * time.Second)
					goto here;
				}
				if len(fs) == 2 {

				} else {
					fmt.Fprintln(conn, "Ambiguous Value")
				}

			case  "LIST":
				if auth == 0 {
					time.Sleep(10 * time.Second)
					goto here;
				}
			case  "NAMES":
				if auth == 0 {
					time.Sleep(10 * time.Second)
					goto here;
				}
			case  "PRIVMSG":
				if auth == 0 {
					time.Sleep(10 * time.Second)
					goto here;
				}
			case  "PART":
				if auth == 0 {
					time.Sleep(10 * time.Second)
					goto here;
				}

			default:
				fmt.Fprintln(conn, "INVALID COMMAND "+fs[0]+"\r\n")
				continue
			}









	}
	defer conn.Close()

	fmt.Println(Uname, "Code Got To The Termination", "Or Exited")
}
