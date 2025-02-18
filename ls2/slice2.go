package main

import "fmt"

func main() {

	var slice1 []string
	slice2 := []string{"abby", "ghost", "samuel"}

	slice1 = append(slice1, "test1")

	fmt.Println(slice1)

	format := 1
	slice2 = append(slice2[:format], slice2[format+1:]...)
	fmt.Println(slice2)

}
