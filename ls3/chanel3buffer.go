package main

import "fmt"

func main() {
	intch := make(chan int, 3) //Создаем канал
	go squaree(intch)          //Запускаем функцию (которая ожидает получение значения в канал)
	intch <- 5                 //Передаем значение
	intch <- 15
	intch <- 50
	fmt.Println(<-intch) //Получаем значение (из функции)
	fmt.Println(<-intch)
	fmt.Println(<-intch)
}

func squaree(ch chan int) {
	num := <-ch     //Получаем значение из канала
	ch <- num * num //Возвращаем новое значение в канал
}
