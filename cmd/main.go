package main

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	score  int
	name   string
	color  string
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn:   conn,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		score:  0,
	}
}

func main() {
	fmt.Println("\033[31mSTARTING SERVER\033[0m")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error from net", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("got new connection from: %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	// for {

	line, err := reader.ReadString('\n')
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Connection timed out")
			conn.Close()
		} else {
			fmt.Println("Client disconnected / conn failed.", err)
		}
	}

	fmt.Printf("Got a line: %q\n", line)

	res := "HTTP/1.1 200 OK\r\nContent-Length: 2\r\n\r\nOK"

	conn.Write([]byte(res))
	defer conn.Close()

	// }
}
