package main

import "fmt"

func main() {
	intCh := make(chan int, 2)
	intCh <- 10
	intCh <- 3
	close(intCh)
	// intCh <- 24          // ошибка - канал уже закрыт
	// fmt.Println(<-intCh) // 10
	// fmt.Println(<-intCh) // 3
	// ---
	// if valss, opened := <-intCh; opened {

	// 	fmt.Println(valss)
	// } else {
	// 	fmt.Println("cnannel closed")
	// }
	// ---
	// if valss, opened := <-intCh; opened {
	// 	fmt.Println(valss)
	// } else {
	// 	fmt.Println("cnannel closed")
	// }
	// ---
	// for i := 0; i <= cap(intCh); i++ {
	// 	if vals, opened := <-intCh; opened {
	// 		fmt.Println(vals)
	// 	} else {
	// 		fmt.Println("channel closed")
	// 	}
	// }
	//alternative
	for {
		val, opened := <-intCh
		if !opened {
			fmt.Println("channel closed")
			break
		}
		fmt.Println(val)
	}
}
