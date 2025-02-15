package main

import "fmt"

func main() {
	var (
		c        = "text"
		e string = "sup"
	)
	d := 3
	var names, ages = "Thom", 22
	const goat = 22
	a := "Golang"
	b := "Very interest"

	fmt.Println(a+b+c+e, names)
	fmt.Println(d*goat + ages)
}
