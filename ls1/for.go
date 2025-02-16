package main

import "fmt"

func main() {

	for i := 1; i <= 9; i++ {
		for a := 1; a <= 9; a++ {
			fmt.Print(i*a, "\t")
		}
		fmt.Println()
	}
}
