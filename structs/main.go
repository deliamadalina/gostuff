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
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName:  "Andreson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "John"
	// alex.lastName = "Smith"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "jsmith@gmail.com",
			zipCode: 026663,
		},
	}

	/*
	   Turn "address" into "value" with "*address"
	   Turn "value" into "address" with "&value"
	*/

	// jimPointer := &jim // memory address of jim/pointer; jim is a value and &jim is the address associated to that value
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy") // the simplified way of the above two lines
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) { // *person - it's a type decision. Here we're working with a poinetr to a person
	/*
	 *pointerToPerson - this is an operator. It means we want to manipulate the value the pointer is referencing
	 */
	(*pointerToPerson).firstName = newFirstName //*pointerToPerson gets the value of the pointer , like the value associated to the address jimPointer
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
