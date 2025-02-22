package main

import "fmt"

type contacts struct {
	phone string
	mail  string
}

type person struct {
	name        string
	age         int
	contactData contacts
}

func (name *person) updateContact(data string, types int) {
	switch types {
	case 1:
		(*name).name = data
	case 2:
		(*name).contactData.phone = data
	case 3:
		(*name).contactData.mail = data
	default:
		fmt.Println("NOT SELECTED")
	}
}

func (list *lister) printall() {
	vals := len(*list)

	for i := vals - 1; i >= 0; i-- {
		fmt.Println((*list)[i])

	}

}

type lister []string

func main() {
	Dmitry := person{name: "Dmitry", age: 21, contactData: contacts{phone: "12345678", mail: "mail@mail.ru"}}
	fmt.Println("----start----")
	fmt.Println(Dmitry)

	Dmitry.updateContact("1488", 2)
	fmt.Println(Dmitry)

	peoples := lister{"1", "2", "3"}
	peoples = append(peoples, "4", "20")
	fmt.Println(peoples)
	peoples.printall()
}
