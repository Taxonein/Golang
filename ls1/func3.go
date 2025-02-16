package main

import "fmt"

func main() {

	var f func(int, int) int = add
	fmt.Println(f(20, 3))

	action(30, 26, multipl)

}
func add(x, y int) int {
	return x + y
}

func multipl(x, y int) int {
	return x * y
}

func action(x, y int, operation func(int, int) int) {
	result := operation(x, y)
	fmt.Println(result)
}
