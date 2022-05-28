package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type xxx struct {
	string
}

func main() {
	test := xxx{"potato"}
	fmt.Printf("%+v", test)

	alex := person{"Alex", "Anderson", contactInfo{"aj@example.com", 12345}}
	fmt.Println(alex)

	bob := person{
		firstName: "Bob",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email: "bsmith@example.com",
			zip:   11111,
		},
	}
	fmt.Println(bob)

	var alice person
	fmt.Println(alice)

	alex.lastName = "Johnson"
	alex.contactInfo.email = "a.j@example.com"
	alex.contactInfo.zip = 90210
	fmt.Printf("%+v", alex)

	alexPointer := &alex // & gives memory address of the variable
	alexPointer.updateName("al")
	alex.print()

	alex.updateName("alexsander")
	alex.print()
}

func (p person) print() {
	fmt.Printf("\nPERSON INFO = %+v\n", p)
}

// *person is a type description it means "a pointer that points at a person"
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointerToPerson - here * is an operator meaning we want to manipulate the value the pointer is referencing
	(*pointerToPerson).firstName = newFirstName
}
