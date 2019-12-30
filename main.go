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
var Global_Users []User
var Global_Channel []ChatChannel

var conn net.Conn
type Connn struct{
	conn []net.Conn

}
type User struct {
	Username string
	Nickname string
	Password string
	Status 	int 
	conn 	net.Conn
	// Channel  ChatChannel
}

	// var Users []User
type ChatChannel struct {
	Name string
	Description string
	Users []User
}

type ChatServer struct {
	// Users []User
	Channels [] ChatChannel
}
type ChatUsers struct {
	Name  string
	Users [] User

}

// type Message struct {
// 	UserClient string // which could be Userclient User
// 	UserMessage string
// }

func main() {
	// var wg sync.WaitGroup

	// usr1 := User{"Femi", "Fem", "0000", 0}
	// usr2 := User{"Victoria", "Ria", "1234", 0}
	tinder := ChatUsers{Name: "Tinder"}	//Chat Users

	channel1 := ChatChannel{Name: "#Welcome", Description: "first"}

	
	server1 := ChatServer{}
	Global_Channel = append(Global_Channel, channel1)

	// tinder.Users = append(tinder.Users, usr1)
	// tinder.Users = append(tinder.Users, usr2)
	// fmt.Println(tinder.Users)
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
		go handleconn(conn, tinder, server1)
	}
	
}

func handleconn(conn net.Conn, tinder ChatUsers, server1 ChatServer) {
	// for {
	// fmt.Println("Conn_ID:", conn)
	// fmt.Fprintf(conn, "Write something: ")
	// scanner := bufio.NewScanner(conn) 
	// scanner.Scan()
	// Nname := scanner.Text()
	// fmt.Fprintf(conn, Nname)
	// }

	fmt.Fprintf(conn, "NOTICE AUTH :*** Looking up your hostname...\nNOTICE AUTH :*** Found your hostname, welcome back\nNOTICE AUTH :*** Checking ident\nNOTICE AUTH :*** No identd (auth) response\n") //displays on the conn client
	
	var Uname , Nname, pass string
	fmt.Println(Uname)// need to store this in my database

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
		"\tNAMES <#channel> \r\n"+				//lists all the Nicknames of users on a Channel 
		"\tNAMES \r\n"+				//lists all the Nicknames of users in a server		
		"\tLIST <#channel> \r\n"+				//lists all channel or current status of channel
		"\tPRIVMSG <nickname>/<channel> \r\n\r\n\r\n\r\n")	//sends a message to another user or channel 

	
		//  memc Connn := append()

		var currentUserID int


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

			for i:=0; i < len(Global_Users); i++  {
				if	Global_Users[i].Username == Uname {
					fmt.Fprintln(conn,"Username Exists, Checking if Password Matches")
					if Global_Users[i].Password == pass {
						fmt.Fprintln(conn,"Welcome Back")
						currentUserID = i
						Global_Users[i].Status = 1
						Global_Users[i].Nickname = Nname
						auth+=1
						time.Sleep(3 * time.Second)
	
						goto here;
					} else {
						fmt.Fprintln(conn,"Password Incorrect")
						auth =0
						time.Sleep(3 * time.Second)
		
						goto here;
					}
					
				}				
			}
			fmt.Fprintln(conn,"Creating Username: ", Uname)
			new_user := User{Username: Uname, Password: pass, Nickname: Nname, Status:1, conn: conn}
			Global_Users = append(Global_Users, new_user)
			tinder.Users = append(tinder.Users, new_user)
			auth+=1
			currentUserID = len(tinder.Users) - 1
			fmt.Println("New User: ", new_user)
			fmt.Println("Whole Array: ", Global_Users)

		case  "NICK":
			if auth == 0 {
				fmt.Fprintln(conn,"You are currently not logged In")
				time.Sleep(3 * time.Second)
	
				goto here;
			}
			if len(fs) == 2 {
				for i:=0; i < len(Global_Users); i++ {
					if Global_Users[i].Username == Uname {
						Global_Users[i].Nickname = fs[1]
						Nname = fs[1]
						fmt.Fprintln(conn, "Nickname Changed To ", Nname)
						break
					}
				}
			} else {
				fmt.Fprintln(conn, "Ambiguous Value")
				fmt.Fprintln(conn, "Nickname not changed")
		}
		case  "JOIN": //JOIN #channel
				if auth == 0 {
					fmt.Fprintln(conn, "Please connect using PASS NICK USER")
					time.Sleep(3 * time.Second)
					goto here;
				}
				if len(fs) == 2 {
					//loop thru list of channels and see is fs(1) == channel
					for i:=0; i < len(Global_Channel); i++ {
						if Global_Channel[i].Name == fs[1]{
							//append user to channel
							// charan thoughts: u have Uname a string. loop thru users in tinder, if tinder.Users[i].Name == Uname, 
							Global_Channel[i].Users = append(Global_Channel[i].Users, Global_Users[currentUserID])
							
							fmt.Fprintln(conn, "Joined the Channel", fs[1])
							goto here;
						} 
					}
					//else, create chatChannel with given name, append user && append chatChannel to chatServer
					fmt.Fprintln(conn, "NOT a Current Channel")

					channel1 := ChatChannel{Name: fs[1], Description: "first"}
					Global_Channel = append(Global_Channel, channel1)
					Global_Channel[len(Global_Channel) - 1].Users = append(Global_Channel[len(Global_Channel) - 1].Users, Global_Users[currentUserID])
					fmt.Fprintln(conn, "Channel Created\nJoined the Channel", fs[1])

					// server1.Channels[len(server1.Channels) - 1].ChatServer = append(server1.Channels[len(server1.Channels) - 1].Users, channel2)
				} else {
					fmt.Fprintln(conn, "Ambiguous Value")
				}

			case  "LIST": // LIST <#channel>
			if auth == 0 {                                
				fmt.Fprintln(conn, "Please connect using PASS NICK USER")
				time.Sleep(3 * time.Second)
				goto here;
			}
			if len(fs) == 2 || len(fs) == 1 {
				for i:=0;i < len(Global_Channel); i++ {
					//printing out all the servers
					fmt.Fprintln(conn,Global_Channel[i].Name)
				}
			} else {
				fmt.Fprintln(conn, "Use the COMMAND <LIST #channel>")
			}

		case  "NAMES":
			fmt.Println("NAMES GLOBAL", Global_Users)

			if auth == 0 {
				fmt.Fprintln(conn, "Please connect using PASS NICK USER")
				time.Sleep(3 * time.Second)
				goto here;
			}
			if len(fs) == 1 {
				fmt.Fprintln(conn, "Listing All The Users Present On the Server")
				// Checking for all the users on the server .... ChatUser
				for i:=0;i < len(Global_Users); i++ {
					//printing out all the servers
					fmt.Fprintln(conn, Global_Users[i].Username)
					
				}
				time.Sleep(3*time.Second)
				goto here;
			}
			if len(fs) == 2 {
				// Checking for all the Users in the Channel
				fmt.Fprintln(conn, "Listing all the User(NICK) presently on the Channel Specified ", fs[1])
				// Check if the Channel Exists
				fmt.Println("fs[1] = |", fs[1], "|")
				for i:=0;i < len(Global_Channel); i++ {
					//checking if the server exists
					if Global_Channel[i].Name == fs[1] {
					//printing out all the servers
						for j:=0; j<len(Global_Channel); j++ {
							fmt.Println("Nickname = ", Global_Channel[i].Users[j].Username)
						}
					} else {
						//Server Doesnot Exist
					fmt.Fprintln(conn,"Channel You Currently Looking does not Exist")
						goto here;
					}
					// fmt.Fprintln(conn, server1.Channels[i])
				}
			} else {
				fmt.Fprintln(conn, "Use the COMMAND <LIST #channel>")
			}
		case  "PRIVMSG":
			if auth == 0 {
				fmt.Fprintln(conn, "Please connect using PASS NICK USER")
				time.Sleep(4 * time.Second)
	
				goto here;
			}
			if len(fs) != 2 {
				fmt.Fprintln(conn, "Ambiguous Value")
				goto here;
			}
			for i:=0;i < len(server1.Channels); i++ {
					if fs[1] == server1.Channels[i].Name{
						// I have found my conn username
					}
			}
		case  "PART":
			if auth == 0 {
				time.Sleep(3 * time.Second)
		
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
