package main

import "fmt"

func main() {

	hello()
	hello()
	hello()
	hello()
	hello()
	multiply(3, 6, 36.6, 23.2)
	wtf(1, 166, 2435, 575)
	test := [...]int{1, 2, 3, 4} //массив
	wtf(test[:]...)
	wtf([]int{1, 2, 3}...) //срез
	fmt.Println(wtf2(2, 8, 2))
}

func hello() {
	fmt.Println("hello")
}
func multiply(a, b int, c, d float32) {
	fmt.Println("int", a*b)
	fmt.Println("float", c+d)
}
func wtf(numbers ...int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum)
}
func wtf2(tips ...int) int {
	sum := 0
	for _, tips := range tips {
		sum += tips
	}
	return (sum)
}
