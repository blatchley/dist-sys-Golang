package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

// GetOutboundIP preferred outbound ip of this machine
// based on code taken from https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go/37382208#37382208
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	hostip, _, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		log.Fatal(err)
	}

	return hostip
}

//sendMessage sends a string to the server encoded using the gob package.
func sendMessage(conn net.Conn, message *string) {
	//Create encode on the connection
	enc := gob.NewEncoder(conn)
	//encode the message
	err := enc.Encode(*message)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

//handleConnection takes a connection, then runs a loop, creating decoders for each message sent.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		//The variable to place the message into
		var msg string
		dec := gob.NewDecoder(conn)
		err := dec.Decode(&msg)
		//Catch decoder errors
		if err != nil {
			fmt.Println("gob decode: " + err.Error())
			return
		}
		//Print our the string
		fmt.Print("received string through bufio reader: " + msg)
	}
}

//Run the "server" functionality.
// delegates new connections to a handleConnection.
func runServer() {
	// leaving the port as ":0" allows go to choose an available port on the machine
	ln, err := net.Listen("tcp", ":0")
	// standard boilerplate for catching errors
	if err != nil {
		log.Fatal(err)
	}

	//get outbound IP address
	ipAddress := GetOutboundIP()
	//get the port the listener is currently listening on
	port := strconv.Itoa(ln.Addr().(*net.TCPAddr).Port)
	ipAndPort := ipAddress + ":" + port
	fmt.Println("Listening for connections on IP:port " + ipAndPort)

	//Loop to accept incoming connections
	for {
		conn, _ := ln.Accept()
		fmt.Println("Got a connection...")
		go handleConnection(conn)
	}
}

func main() {
	hostAddress := ""
	fmt.Println("Please enter server IP and port (in the form hostIP:port) to connect, or type \"new\" to start a new server")
	fmt.Scanln(&hostAddress)
	//Code for han
	if hostAddress == "new" {
		runServer()
	} else {
		//Connect to the specified server
		conn, err := net.Dial("tcp", hostAddress)
		if err != nil {
			log.Fatal(err)
		}

		//Create reader to take in new messages to send to server.
		reader := bufio.NewReader(os.Stdin)

		//Loop reading lines from stdin, and sending them.
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("User input loop: " + err.Error())
			}
			sendMessage(conn, &text)
		}

	}
}
