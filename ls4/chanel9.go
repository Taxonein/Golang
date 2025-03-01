package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	work := func(id int) {
		fmt.Println("Горутина ", id, " начала выполнение")
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина ", id, " закончила выполнение")
		defer wg.Done()
	}

	go work(1)
	go work(2)

	wg.Wait()

}
