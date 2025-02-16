package main

import "fmt"

func main() {

	colors := [3]string{2: "blue", 0: "red", 1: "gog"}

	fmt.Println(colors[0])

	if colors[2] == "blue" {
		fmt.Println("TRUE")
	} else if colors[1] == "gog" {
		fmt.Println("goga")
	} else {
		fmt.Println("test")
	}
}
