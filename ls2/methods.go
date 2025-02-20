package main

import "fmt"

type person struct {
	name string
	age  int
}

func (a person) print() {
	fmt.Println(a.age)
	fmt.Println(a.name)
}

func (a person) eat(meal string) {
	fmt.Println(a.name + " eст " + meal)

}

func main() {
	var tom = person{name: "Tom", age: 19}
	tom.print()
	tom.eat("Борщ")

	tomP := &tom
}
