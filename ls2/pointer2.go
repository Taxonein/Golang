package main

import "fmt"

func mat(x *int) {
	*x = (*x) * (*x)
}

func main() {
	d := 5
	fmt.Println(d)
	mat(&d)
	fmt.Println(d)
}
