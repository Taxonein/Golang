package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
func megaadd(numbs ...int) int {
	result := 0
	for _, num := range numbs {
		fmt.Println(num)
		result += num
	}
	return result
}

func mathe(a, b int, op func(int, int) int) int {
	res := op(a, b)
	return res
}

func main() {

	fmt.Println(mathe(10, 15, add))
	fmt.Println(megaadd(30, 55, 222, 623, 7))

}
