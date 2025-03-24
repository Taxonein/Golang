package main

import (
	"fmt"
	"net"
)

func main() {
	msg := "hello from server!"

	listener, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("listening...")
	for {
		con, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		con.Write([]byte(msg))
		con.Close()
	}
}
