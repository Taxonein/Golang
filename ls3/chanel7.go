package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {
	ch := make(chan bool)
	var mutex sync.Mutex

	for i := 0; i <= 5; i++ {
		go gol(i, ch, &mutex)
	}

	for i := 1; i <= 5; i++ {
		<-ch
	}

}

func gol(i int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", i, "-", counter)
	}
	mutex.Unlock()
	ch <- true
}
