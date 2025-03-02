package main

import (
	"errors"
	"fmt"
	"io"
)

type phone string

func (ph phone) Read() (string, error) {
	var result string
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			result += string(ph[i])
		}
	}
	if result == "" {
		return "", errors.New("Not enought digits")
	}
	return result, io.EOF

}

func main() {
	var num1 phone = "+7(903-532-76-44)"
	exit, err := num1.Read()
	fmt.Println(exit)
	fmt.Println(err)

}
