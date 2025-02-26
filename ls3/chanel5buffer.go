package main

import "fmt"

var inch chan<- int  //Отправка данных
var outch <-chan int //Получение данных

func main() {

	intch := make(chan int, 2)

	//	go mult(5, intch)
	//	go mult(10, intch)

	//	fmt.Println(<-intch)
	//	fmt.Println(<-intch)
	intch <- 5
	fmt.Println("1")
	if val, succ := <-intch; succ {
		fmt.Println(val)
		fmt.Println(succ)
	}

}

func mult(num int, ch chan<- int) {
	ch <- num * num
}
