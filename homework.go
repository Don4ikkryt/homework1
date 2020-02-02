package main

import (
	"fmt"
)

type Work interface {
	printInformaionAboutWork()
}
type Worker interface {
	printInformationAboutWorker()
}

//person
type person struct {
	age         int
	gender      string
	sitizenShip string
	lastName    string
	firstName   string
}

func newPerson(age int, gender, sitizenShip, lastName, firstName string) person {
	return person{age: age, gender: gender, sitizenShip: sitizenShip, lastName: lastName, firstName: firstName}
}
func (p person) printInformationAboutPerson() {
	fmt.Println("Person: ")
	fmt.Print("Firstname: ")
	fmt.Println(p.firstName)
	fmt.Print("Lastname: ")
	fmt.Println(p.lastName)
	fmt.Print("Age: ")
	fmt.Println(p.age)
	fmt.Print("Gender :")
	fmt.Println(p.gender)
	fmt.Print("Sitizenship: ")
	fmt.Println(p.sitizenShip)
}

//person

//programming
type programming struct {
	_programmer programmer
}
type programmer struct {
	_person person
}

//programming

//bank
type bank struct {
	_banker banker
}
type banker struct {
	_person person
}

//bank

//taxi
type taxi struct {
	_taxiDriver taxiDriver
}
type taxiDriver struct {
}

//taxi

//school
type school struct {
	_teacher teacher
}
type teacher struct {
	_person person
}

//school

//factory
type factory struct {
	_workerOfFactory workerOfFactory
}

type workerOfFactory struct {
	_person person
}

//factory
func main() {

}
