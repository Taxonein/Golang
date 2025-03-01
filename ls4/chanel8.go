package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {

	ch := make(chan bool)
	for i := 0; i <= 5; i++ {
		go doit(i, ch, &mutex)
	}

	for a := 0; a <= 5; a++ {
		<-ch
	}

}
func doit(num int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	for i := 0; i <= 5; i++ {
		fmt.Println(num, " has ", i)
	}
	mutex.Unlock()
	ch <- true

}
