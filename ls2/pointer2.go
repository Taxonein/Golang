package main

import "fmt"

func mat(x *int) {
	*x = (*x) * (*x)
}

func main() {
	d := 0
	fmt.Println(d)
	mat(&d)
	fmt.Println(d)
	if &d == nil {
		fmt.Println(&d)
	}
}
