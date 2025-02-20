package main

import "fmt"

type library []string

func (a library) print() {
	for _, val := range a {
		fmt.Println(val)
	}
}

func main() {

	var lib library = library{"book1", "book2", "book3"}
	lib.print()

}
