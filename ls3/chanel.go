package main

import "fmt"

func factoriall(num int, ch chan int) {
	if num < 1 {
		panic("Number is under 1")
	}
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	fmt.Println(num, "-", result)
	ch <- result
}

func main() {

	intch := make(chan int)

	go factoriall(5, intch)
	fmt.Println(<-intch)
	//fmt.Println(<-intch)
}
