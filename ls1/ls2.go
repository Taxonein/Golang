package main

import "fmt"

func main() {
	var (
		name string = "Goga"
		age int = 20
	)

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(name + " text " + fmt.Sprint(age))
}