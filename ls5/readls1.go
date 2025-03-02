package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	data := "Hello, Golang!"

	reader := strings.NewReader(data)

	buf := make([]byte, len(data))

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Достигнут лимит строки")
				break
			}
		}
		fmt.Println("Ошибка чтения")
		fmt.Println("Всего ", n, "байт:", string(buf[:n]))
		break
	}

}
