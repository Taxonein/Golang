package main

import "fmt"

func main() {
	intch := make(chan int) //Создаем канал
	go square(intch)        //Запускаем функцию (которая ожидает получение значения в канал)
	intch <- 5              //Передаем значение
	fmt.Println(<-intch)    //Получаем значение (из функции)
}

func square(ch chan int) {
	num := <-ch     //Получаем значение из канала
	ch <- num * num //Возвращаем новое значение в канал
}
