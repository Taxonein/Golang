package main

import "fmt"

func main() {
	defer exec()
	fmt.Println("test1")

}

func exec() {
	fmt.Println("test")
}
