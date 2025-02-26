package main

import "fmt"

func factorial(num int) {
	if num < 1 {
		panic("Number is under 1")
	}
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	fmt.Println(num, "-", result)
}

func main() {
	for i := 1; i <= 7; i++ {
		go factorial(i)
	}
	fmt.Scanln()
	fmt.Println("end")
}
