package main

import "fmt"

// source
type PersonIntf interface {
	FirstName() string
	LastName() string
}

type Person struct {
	firstname, lastname string
}

func (p Person) FirstName() string {
	return p.firstname
}

func (p Person) LastName() string {
	return p.lastname
}

func NewPerson(firstname, lastname string) *Person {
	return &Person{firstname: firstname, lastname: lastname}
}

// target
type EmployeeIntf interface {
	FullName() string
}

type Employee struct {
	fullname string
}

func (e Employee) FullName() string {
	return e.fullname
}

// usage
func PrintEmployeeName(e EmployeeIntf) {
	fmt.Println(e.FullName())
}

// adapter
type PersonEmployeeAdapter struct {
	person PersonIntf
}

func (a PersonEmployeeAdapter) FullName() string {
	return fmt.Sprintf("%s %s", a.person.FirstName(), a.person.LastName())
}

func NewPersonEmployeeAdapter(p PersonIntf) *PersonEmployeeAdapter {
	return &PersonEmployeeAdapter{person: p}
}

func main() {
	p := NewPerson("John", "Smith")
	adapter := NewPersonEmployeeAdapter(p)
	PrintEmployeeName(adapter)
}
