package main

import "fmt"

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	var numbers2 [5]int = [5]int{1, 5}
	numbers3 := [3]int{2, 4, 6}
	numbers4 := [...]int{1, 2, 3, 4}

	fmt.Println(numbers)
	fmt.Println(numbers2)
	fmt.Println(numbers3)
	fmt.Println(numbers4)
	fmt.Println(numbers2[1])

}
