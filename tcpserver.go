package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var incufridgeConn net.Conn
var incufridgeReader *bufio.Reader

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	// will listen for message to process ending in newline (\n)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Client disconnected")
		return
	}

	name = strings.TrimSpace(name)

	log.Println(name + " has connected")

	if name == "incufridge" {
		incufridgeConn = conn
		incufridgeReader = bufio.NewReader(conn)
		conn.Write([]byte("received\n"))
	} else if name == "client" {
		conn.Write([]byte("received\n"))
		for {
			msg, err := reader.ReadString('\n')
			msg = strings.TrimSpace(msg)
			log.Println("received message: " + msg)
			if err != nil {
				log.Println("client disconnected")
				return
			}
			incufridgeConn.Write([]byte(msg + "\n"))
			log.Println("Sent message")
			response, _ := incufridgeReader.ReadString('\n')
			response = strings.TrimSpace(response)
			log.Println("Received response: ", response)
			conn.Write([]byte(response + "\n"))
		}
	}
	// output message received
	// log.Print("Message Received:", string(message))
	// sample process for string received
	// newmessage := strings.ToUpper(message)
	// send new string back to client
	// conn.Write([]byte(newmessage + "\n"))
}

func main() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":26517")

	for {
		conn, _ := ln.Accept()
		go handleClient(conn)
	}
}
