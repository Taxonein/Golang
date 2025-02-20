package main

import "fmt"

type multe = int

func action(a, b int, oper func(int, int) int) {
	result := oper(a, b)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	c := a * b
	return (c)
}

func mathe(a, b int, funct func(int, int) int) {
	mathes := funct(a, b)
	fmt.Println(mathes)

}

func main() {

	action(20, 30, add)
	mathe(30, 30, mult)

}
