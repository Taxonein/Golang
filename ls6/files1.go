package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		panic("err")
	}
	fmt.Println(file.Name())

	filee, errorr := os.Open("hello.txt")
	if errorr != nil {
		panic("err")
	}
	filee.Close()

}
