package main

import "fmt"

func main() {

	f := square()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	test := func(x, y int) int { return x * y }

	fmt.Println(test(8, 8))

}

func square() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}

func squares() func() int {
	var num int = 6
	return func() int {
		num++
		return num
	}
}
