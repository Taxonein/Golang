package main

import (
	"fmt"
	"io"
	"strings"
)

type strok string

func (data strok) Read() {
	reader := strings.NewReader(string(data))
	byt := make([]byte, len(data))

	for {
		n, err := reader.Read(byt)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Лимит строки")
				break
			}
			fmt.Println("Закончено")
			break
		}
		fmt.Println("Всего имеется ", n, " байт. Строка: ", string(byt))
	}
}

func main() {

	zamn := strok("nice cock")
	zamn.Read()

	gadd := strok("+7903143121")
	gadd.Read()

}
