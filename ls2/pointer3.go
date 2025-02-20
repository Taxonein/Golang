package main

import "fmt"

func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}

func main() {

	a := 6

	b := &a

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)

	//s := createPointer(20)
	//fmt.Println(*s)

}
