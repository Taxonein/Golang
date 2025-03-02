package main

import (
	"fmt"
	"io"
)

type phreader string

func (ph phreader) Read(p []byte) (int, error) {

	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	return count, io.EOF

}

func main() {
	phone1 := phreader("+7-903-531-36-66")
	buffer := make([]byte, len(phone1))

	phone1.Read(buffer)
	fmt.Println(phreader(buffer))
}
