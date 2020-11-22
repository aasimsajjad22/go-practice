package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//p := person{firstName: "Aasim", lastName: "Sajjad"}

	p := person{
		firstName: "Aasim",
		lastName:  "Sajjad",
		contactInfo: contactInfo{
			email:   "test@gmail.com",
			zipCode: 44000,
		},
	}

	p.updateName("Simm")
	p.print()

}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
