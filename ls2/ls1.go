package main

import "fmt"

func main() {
	defer exec()
	fmt.Println("test1")

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		if i == 99 {
			panic(fmt.Sprint("Error, used value is: ", i))
		}
	}

}

func exec() {
	fmt.Println("test")
}
func logger() {
	fmt.Println("You have error!")
}
