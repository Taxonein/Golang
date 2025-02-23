package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)

	fmt.Println("Say you name: ")
	fmt.Scanf("%s\n", &name)

	fmt.Println("And your age: ")
	fmt.Scanf("%d\n", &age)

	fmt.Println("Hello, Your age is" + name)
}
