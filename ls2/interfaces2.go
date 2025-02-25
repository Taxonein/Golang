package main

import "fmt"

type stream interface {
	read()
	write()
	close()
}
type file struct {
	name string
	text string
}
type folder struct{}

func (f file) doit() {
	fmt.Println(f.name)
}

func main() {
	file.doit(file{})
	testt := file{name: "1233", text: "1222"}
	testt.doit()

}
