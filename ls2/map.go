package main

import "fmt"

func main() {

	var people = map[string]int{
		"Tom":  1,
		"Bob":  2,
		"Dima": 3,
	}

	if val, ok := people["Tom"]; ok { //Если
		fmt.Println(val)
		fmt.Println(ok)
	}

}
