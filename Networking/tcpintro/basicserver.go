package main

import (
	"fmt"
	"log"
	"net"
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

// handleconnection receives incoming connections, prints their ip, then closes them.
func handleConnection(conn net.Conn) {
	fmt.Println("Received a connection from " + conn.RemoteAddr().String())
	fmt.Println("Closing connection again.")
	conn.Close()
}

//Run the "server" functionality
func runServer() {
	// leaving the port as ":0" allows go to choose an available port on the machine
	ln, err := net.Listen("tcp", ":0")
	// standard boilerplate for catching errors
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	//get outbound IP address
	ipAddress := GetOutboundIP()
	//get the port the listener is currently listening on
	port := strconv.Itoa(ln.Addr().(*net.TCPAddr).Port)
	ipAndPort := ipAddress + ":" + port
	fmt.Printf("Listening for connections on IP:port " + ipAndPort)

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
	//run the server
	if hostAddress == "new" {
		runServer()

	} else {
		//Connect to the server
		conn, err := net.Dial("tcp", hostAddress)
		if err != nil {
			log.Fatal(err)
		}

		//Close the connection
		conn.Close()
	}
}
