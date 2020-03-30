package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	name := conn.RemoteAddr().String()

	fmt.Printf("%+v connected\n", name)

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println("Request: ", text)
		if text == "" {
			break
		}
	}
	fmt.Println("Write anser")
	fmt.Fprintln(conn, "HTTP/1.1 200 OK")
	fmt.Fprintln(conn, "Server: golag-net")
	fmt.Fprintln(conn, "Set-Cookie: specific=value")
	fmt.Fprintln(conn)
	fmt.Fprintln(conn, "<h1>Hello</h1>")
	fmt.Fprintln(conn)
	conn.Write([]byte("Hello, " + name + "\n\r"))
}

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
