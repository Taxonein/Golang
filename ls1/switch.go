package main

import "fmt"

func main() {
	a := 66

	switch a {
	case 8:
		fmt.Println("8")
	case 9:
		fmt.Println("9")
	case 10:
		fmt.Println("10")
	case 11, 12, 13:
		fmt.Println("super")
	default:
		fmt.Println("gogsss")
	}
}
