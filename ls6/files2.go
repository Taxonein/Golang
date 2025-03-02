package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file1, err := os.OpenFile("hello.txt", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	_, err = file1.WriteString("ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!ПРИВЕТ!")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	file1.Close()

	// Читаем файл

	file2, err := os.OpenFile("hello.txt", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	data := make([]byte, 64)
	var result []byte

	for {
		n, err := file2.Read(data)
		if err != nil && err == io.EOF {
			break
		}
		result = append(result, data[:n]...)
	}
	fmt.Println(string(result))
	file2.Close()
}
