package main

import "fmt"

type person struct {
	name string
	age  int
}

func (guy *person) updateAge(newAge int) {
	(*guy).age = newAge
}

func main() {

	thomas := person{name: "Thomas", age: 21}
	var thomasPoint *person = &thomas

	fmt.Println(thomasPoint)

	fmt.Println(thomas)
	thomas.updateAge(22)
	fmt.Println(thomas)
}
