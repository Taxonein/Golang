package main

import "fmt"

func main() {

	fmt.Println(giveformat(21, "Dmitry", "Yan"))
	fmt.Println(giveformat(76, "Vasiliy", "Yan"))

}
func giveformat(age int, firstname, lastname string) (int, string) {
	ageform := age - 18
	nameform := (lastname + " " + firstname)
	return ageform, nameform
}
func giveformat2(age int, firstname, lastname string) (ageform int, nameform string) {
	ageform = age - 18
	nameform = (lastname + " " + firstname)
	return
}
