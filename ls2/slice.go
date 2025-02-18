package main

import "fmt"

func main() {

	users := []string{"Alpha", "Beta", "Gamma", "Delta"}

	users[0] = "Alpen"

	var users3 []string = make([]string, 3) //Создает срез из трех элементов

	var users4 = []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}

	users3 = append(users3, "Boba")

	// for index, value := range users {

	// 	fmt.Println(index, value)
	// }

	var le = 2
	var test []string
	test = append(users4[:le], users4[le+1:]...)

	fmt.Println(test)
}
