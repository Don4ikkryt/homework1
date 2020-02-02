package main

import (
	"fmt"
)

type Work interface {
	printInformationAboutWork()
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

	fmt.Print("Firstname: ")
	fmt.Println(p.firstName)
	fmt.Print("Lastname: ")
	fmt.Println(p.lastName)
	fmt.Print("Age: ")
	fmt.Println(p.age)
	fmt.Print("Gender: ")
	fmt.Println(p.gender)
	fmt.Print("Sitizenship: ")
	fmt.Println(p.sitizenShip)
}

//person

//programming
type programming struct {
	_programmer                 *programmer
	quantityOfProgrammsPerMonth int
	salary                      int
}

func (p *programming) hireWorker(_programmer *programmer) {
	p._programmer = _programmer
}
func newProgrammerVacancy(salary int, quantityOfProgrammsPerMonth int) programming {
	return programming{salary: salary, quantityOfProgrammsPerMonth: quantityOfProgrammsPerMonth}
}
func (p programming) printInformationAboutWork() {
	fmt.Println("Programming: ")
	p._programmer.printInformationAboutWorker()
	fmt.Println("Quantity of programms per month: ")
	fmt.Println(p.quantityOfProgrammsPerMonth)
}

type programmer struct {
	_person                        person
	quantityOfProgrammingLangueges int
}

func (p programmer) printInformationAboutWorker() {
	fmt.Println("Programmer: ")
	p._person.printInformationAboutPerson()
	fmt.Println("Quantity of programming langueges: ")
	fmt.Println(p.quantityOfProgrammingLangueges)
}
func newProgrammer(age int, gender, sitizenShip, lastName, firstName string, quantityOfProgrammingLangueges int) programmer {
	return programmer{_person: newPerson(age, gender, sitizenShip, lastName, firstName), quantityOfProgrammingLangueges: quantityOfProgrammingLangueges}
}

//programming

//bank
type bank struct {
	_banker           *banker
	hoursToWorkPerDay int
	salary            int
}

func (b *bank) hireWorker(_banker *banker) {
	b._banker = _banker
}
func newBankerVacancy(salary int, hoursToWorkPerDay int) bank {
	return bank{salary: salary, hoursToWorkPerDay: hoursToWorkPerDay}
}

func (b bank) printInformationAboutWork() {
	fmt.Println("Bank: ")
	b._banker.printInformationAboutWorker()
	fmt.Println("Hours to work per day: ")
	fmt.Println(b.hoursToWorkPerDay)
}

type banker struct {
	_person           person
	efficiencyPerHour int
}

func (b banker) printInformationAboutWorker() {
	fmt.Println("Banker: ")
	b._person.printInformationAboutPerson()
	fmt.Println("Efficiency per hour: ")
	fmt.Println(b.efficiencyPerHour)
}
func newBanker(age int, gender, sitizenShip, lastName, firstName string, efficiencyPerHour int) banker {
	return banker{_person: newPerson(age, gender, sitizenShip, lastName, firstName), efficiencyPerHour: efficiencyPerHour}
}

//bank

//taxi
type taxi struct {
	_taxiDriver           *taxiDriver
	quantityOfRidesPerDay int
	salary                int
}

func (t *taxi) hireWorker(_taxiDriver *taxiDriver) {
	t._taxiDriver = _taxiDriver
}
func newTaxiDriverVacancy(salary int, quantityOfRidesPerDay int) taxi {
	return taxi{salary: salary, quantityOfRidesPerDay: quantityOfRidesPerDay}
}

func (t taxi) printInformationAboutWork() {
	fmt.Println("Taxi: ")
	t._taxiDriver.printInformationAboutWorker()
	fmt.Println("Quantity of rides per day: ")
	fmt.Println(t.quantityOfRidesPerDay)
}

type taxiDriver struct {
	_person person
	car     string
}

func (td taxiDriver) printInformationAboutWorker() {
	fmt.Println("Taxi driver: ")
	td._person.printInformationAboutPerson()
	fmt.Println("Model of car: ")
	fmt.Println(td.car)
}
func newTaxiDriver(age int, gender, sitizenShip, lastName, firstName string, car string) taxiDriver {
	return taxiDriver{_person: newPerson(age, gender, sitizenShip, lastName, firstName), car: car}
}

//taxi

//school
type school struct {
	_teacher               *teacher
	quantityOfPupilsPerDay int
	salary                 int
}

func (s *school) hireWorker(_teacher *teacher) {
	s._teacher = _teacher
}
func newTeacherVacancy(salary int, quantityOfPupilsPerDay int) school {
	return school{salary: salary, quantityOfPupilsPerDay: quantityOfPupilsPerDay}
}

func (s school) printInformationAboutWork() {
	fmt.Println("School: ")
	s._teacher.printInformationAboutWorker()
	fmt.Println("Quantity of pupils per day: ")
	fmt.Println(s.quantityOfPupilsPerDay)
}

type teacher struct {
	_person            person
	quantityOfSubjects int
}

func (t teacher) printInformationAboutWorker() {
	fmt.Println("Teacher: ")
	t._person.printInformationAboutPerson()
	fmt.Println("Quantity of subjects subjects: ")
	fmt.Println(t.quantityOfSubjects)
}
func newTeacher(age int, gender, sitizenShip, lastName, firstName string, quantityOfSubjects int) teacher {
	return teacher{_person: newPerson(age, gender, sitizenShip, lastName, firstName), quantityOfSubjects: quantityOfSubjects}
}

//school

//factory
type factory struct {
	_workerOfFactory      *workerOfFactory
	quantityOfWorkPerHour int
	salary                int
}

func (f *factory) hireWorker(_workerOfFactory *workerOfFactory) {
	f._workerOfFactory = _workerOfFactory
}
func newFactoryWorkerVacancy(salary int, quantityOfWorkPerHour int) factory {
	return factory{salary: salary, quantityOfWorkPerHour: quantityOfWorkPerHour}
}

func (f factory) printInformationAboutWork() {
	fmt.Println("School: ")
	f._workerOfFactory.printInformationAboutWorker()
	fmt.Println("Quantity of work per day: ")
	fmt.Println(f.quantityOfWorkPerHour)
}

type workerOfFactory struct {
	_person           person
	efficiencyPerHour int
}

func (wf workerOfFactory) printInformationAboutWorker() {
	fmt.Println("Factory worker: ")
	wf._person.printInformationAboutPerson()
	fmt.Println("Efficiency per hour: ")
	fmt.Println(wf.efficiencyPerHour)
}
func newWorkerOfFactory(age int, gender, sitizenShip, lastName, firstName string, efficiencyPerHour int) workerOfFactory {
	return workerOfFactory{_person: newPerson(age, gender, sitizenShip, lastName, firstName), efficiencyPerHour: efficiencyPerHour}
}

//factory
func main() {
	factory := newFactoryWorkerVacancy(1000, 5)
	school := newTeacherVacancy(1200, 10)
	taxi := newTaxiDriverVacancy(800, 10)
	programming := newProgrammerVacancy(1500, 4)
	bank := newBankerVacancy(800, 6)
	factoryWorker := newWorkerOfFactory(20, "Male", "Ukrainian", "Denis", "Sydorenko", 7)
	teacher := newTeacher(20, "Male", "Ukrainian", "Denis", "Sydorenko", 7)
	banker := newBanker(20, "Male", "Ukrainian", "Denis", "Sydorenko", 7)
	programmer := newProgrammer(20, "Male", "Ukrainian", "Denis", "Sydorenko", 7)
	taxiDriver := newTaxiDriver(20, "Male", "Ukrainian", "Denis", "Sydorenko", "Mazda")
	factory.hireWorker(&factoryWorker)
	school.hireWorker(&teacher)
	bank.hireWorker(&banker)
	programming.hireWorker(&programmer)
	taxi.hireWorker(&taxiDriver)
	factory.printInformationAboutWork()
	fmt.Println()
	school.printInformationAboutWork()
	fmt.Println()
	bank.printInformationAboutWork()
	fmt.Println()
	programming.printInformationAboutWork()
	fmt.Println()
	taxi.printInformationAboutWork()
}
