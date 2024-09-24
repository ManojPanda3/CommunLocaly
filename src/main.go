package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var user string = "Manoj"
	var ListenForConn string

	fmt.Printf("Enter what you want to do (listen/connect): ")
	fmt.Scanln(&ListenForConn)
	if ListenForConn == "listen" {
		listenForConn(user)
	} else if ListenForConn == "connect" {
		connectToTheServer()
	}
}

func listenForConn(user string) {
	ln, errLn := net.Listen("tcp", ":8080")
	if errLn != nil {
		fmt.Println("Error: ", errLn)
		return
	}
	defer ln.Close()

	fmt.Println("Server is waiting for a connection at :8080")
	conn, errConn := ln.Accept()
	if errConn != nil {
		fmt.Println("Error: ", errConn)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected successfully")
	go handleMsg(conn, "receiving")

	// Use fmt.Fprintf for better formatting
	fmt.Fprintf(conn, "Hello, I am %s\n", user)
	fmt.Print("Type your message: ")
	handleMsg(conn, "sending")
	return
}

func connectToTheServer() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the server.")
	go handleMsg(conn, "receiving")

	fmt.Print("Type your message: ")
	handleMsg(conn, "sending")
	return
}

func handleMsg(conn net.Conn, msgType string) {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	if msgType == "sending" {
		for {
			scanner.Scan()
			msg := scanner.Text()

			if strings.TrimSpace(msg) == "exit" {
				fmt.Println("Exiting...")
				return
			}

			fmt.Fprintf(writer, "%s\n", msg)
			writer.Flush() 

			fmt.Print("Type your message: ")
		}
	} else if msgType == "receiving" {
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error receiving message:", err)
				return
			}
			fmt.Printf("\nServer: %s", msg)

			fmt.Print("Type your message: ")
		}
	}
}
