package main

import "fmt"

func main() {

	users := [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}

	for index, value := range users {
		fmt.Println(index, value)
	}

	for _, value := range users {
		fmt.Println(value)
		if value == "Gamma" {
			fmt.Println("ALARM")
			continue
		}
		if value == "Epsilon" {
			fmt.Println("ALARM")
			break
		}
	}

}
