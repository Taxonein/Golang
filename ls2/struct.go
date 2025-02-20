package main

import "fmt"

func main() {

	type person struct {
		name string
		age  int
	}

	var tom person = person{"Michael", 22}
	bob := person{name: "Bob", age: 31}

	fmt.Println(person(tom))
	fmt.Println(bob.name)

}
