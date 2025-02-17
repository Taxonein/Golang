package main

import "fmt"

func main() {

	act(10, 5, func(x, y int) int { return x * y })

}

func act(a, b int, operation func(int, int) int) {
	res := operation(a, b)
	fmt.Println(res)
}
