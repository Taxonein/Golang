package main

import "fmt"

func main() {

	var people = map[string]int{
		"Tom":  1,
		"Bob":  2,
		"Dima": 3,
	}

	people1 := make(map[string]int)
	people1["Kate"] = 128
	delete(people1, "Kate")

	fmt.Println(people1)

	if val, ok := people["Tom"]; ok { //Если
		fmt.Println(val)
		fmt.Println(ok)
	}
	for key, value := range people {
		fmt.Println(key, value)
	}

}
