package main

import "fmt"

func main() {

	tes := selectfn(2)
	fmt.Println(tes(20, 80))

}

func selectfn(a int) func(int, int) int {
	if a == 1 {
		return func(x, y int) int { return (x * y) }
	} else {
		return func(x, y int) int { return (x + y) }
	}
}
