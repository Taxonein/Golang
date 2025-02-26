package main

import "fmt"

type file3 struct {
	text string
}
type folder3 struct {
	name string
}

type stream3 interface {
	read() string
	write(string)
	close()
}

// *******

func writeToStream(streamm stream3, text string) {
	streamm.write(text)
}
func closeStream(streamm stream3) {
	streamm.close()
}

// *******

func (f *file3) read() string {
	return f.text
}
func (f *file3) write(message string) {
	f.text = message
}
func (f *file3) close() {
	fmt.Println("Файл закрыт: " + f.text)
}

//*******

func main() {
	myfile := file3{}
	//myfile.write("123")
	//fmt.Println(myfile.read())
	//myfile.close()

	writeToStream(&myfile, "YESS")
	closeStream(&myfile)
}
