package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	ceo := person{
		firstName: "Jim",
		lastName:  "Jones",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94009,
		},
	}
	ceo.print()
	ceoPointer := &ceo
	ceoPointer.updateName("John", "Jones")
	ceo.print()
}

func (p person) print() {
	str := "First Name: " + p.firstName + "\n"
	str = str + "Last Name: " + p.lastName + "\n"
	str = str + "\tContact Info:\n"
	str = str + "\t  Email: " + p.contact.email + "\n"
	str = str + "\t  Zip Code: " + fmt.Sprint(p.contact.zipCode) + "\n"

	fmt.Println(str)
}

func (p *person) updateName(firstname string, lastname string) {
	(*p).firstName = firstname
	(*p).lastName = lastname
}