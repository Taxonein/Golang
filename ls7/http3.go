package main

import (
	"io"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != err {
		panic(err)
	}
	io.Copy(os.Stdout, con)

}
