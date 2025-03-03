package main

import (
	"fmt"
	"io"
	"os"
)

type filename string

var file1 filename = "test.txt"

func main() {
	//text := "Abracadabra"
	//texstbytes := []byte(text)
	//file1.Write(texstbytes)

	file1.Read()
}

func (name filename) Write(data []byte) {
	file, err := os.OpenFile(string(name), os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	file.Close()

}

func (name filename) Read() {
	text := make([]byte, 64)
	var result []byte
	file, err := os.OpenFile(string(name), os.O_RDONLY, 0700)
	if err != nil {
		panic(err)
	}
	for {
		n, err := file.Read(text)
		if err != nil && err == io.EOF {
			break
		}
		result = append(result, text[:n]...)
	}
	//file.Seek(0, io.SeekStart) //Сбрасывает указатель текста в начало
	//io.Copy(os.Stdout, file)   // Копирует вывод (содержимое файла) в stdout (консоль)
	fmt.Println(string(result))
	file.Close()

}
