package main

import "fmt"

func main() {
	strch1 := make(chan int, 3)
	go doit(strch1)
	strch1 <- 5
	strch1 <- 10
	strch1 <- 40

	fmt.Println(<-strch1)
	fmt.Println(<-strch1)
	fmt.Println(<-strch1)
}

func doit(veit chan int) {
	star := <-veit
	veit <- star + star
}
