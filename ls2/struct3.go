package main

import "fmt"

type pers struct {
	name string
	age  int
}
type contact struct {
	name    string
	email   string
	contact pers
}

var tom = contact{
	name:  "Gosling",
	email: "test@mail.com",
	contact: pers{
		name: "Nmae",
		age:  12,
	},
}

func main() {

	fmt.Println(tom)

}
