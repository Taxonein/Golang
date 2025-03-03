package main

import (
	"fmt"
	"io"
	"os"
)

type phoneReader string

func (p phoneReader) Read(bs []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			bs[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	phone1 := phoneReader("+1(234)567 90-10")
	phone2 := phoneReader("+7(933)252 90-11")
	fmt.Println(phone2)
	io.Copy(os.Stdout, phone1)
}
