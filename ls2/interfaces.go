package main

import "fmt"

type car struct{ name string }
type aircraft struct{}

type vehicle interface {
	move()
	kill()
}

func (c car) move() {
	fmt.Println("car едет")
}
func (c car) kill() {
	fmt.Println("car убит", c.name)
}

func (a aircraft) move() {
	fmt.Println("aircraft летит")
}

func main() {

	var tesla vehicle = car{"tesla"}
	mazda := car{}
	mazda.move()

	tesla.kill()

}
