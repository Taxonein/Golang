package main

import "fmt"

func main() {

	gov := selectfn(1)

	fmt.Println(gov(2, 3))

}
func calcs(x, y int) int     { return x + y }
func down(x, y int) int      { return x - y }
func multiplyy(x, y int) int { return x * y }
func divide(x, y int) int    { return x / y }

func selectfn(a int) func(int, int) int {
	if a == 1 {
		return calcs
	} else if a == 2 {
		return down
	} else if a == 3 {
		return multiplyy
	} else {
		return divide
	}
}
