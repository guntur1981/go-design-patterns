# Description

The Adapter Design Pattern is a struct which adapts an existing interface X to conform to the required interface Y.

## Example

Consider the following example:

```
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

func main() {
	p := NewPerson("John", "Smith")
	PrintEmployeeName(p) // error, p does not implement EmployeeIntf
}


```

In the example above, we simulate the use of a target interface (Employee) where we only have a source interface (Person). And as expected, an error will appear because the Person interface is not compatible with the Employee interface.

Suppose the Employee interface is in an external package that we cannot modify, we may consider modifying the Person interface, as follows:

```
// source
type PersonIntf interface {
	FirstName() string
	LastName() string
	FullName() string // new
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

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstname, p.lastname)
}
```

## The Reason Why This Is Not a Good Practice

By doing above modification, we have violated the Open-Closed Principle. Not to mention that if the Person interface is also in an external package which we cannot modify it.

## A Better Approach

**First**, create an adapter struct that takes Person interface as an internal property.

```
// adapter
type PersonEmployeeAdapter struct {
	person PersonIntf
}

func NewPersonEmployeeAdapter(p PersonIntf) *PersonEmployeeAdapter {
	return &PersonEmployeeAdapter{person: p}
}
```

**Second**, implement all functions according to the target interface (Employee). From the first example above, only FullName() needs to be implemented.

```
func (a PersonEmployeeAdapter) FullName() string {
	return fmt.Sprintf("%s %s", a.person.FirstName(), a.person.LastName())
}
```

**Finally**, we can use the adapter.

```
func main() {
	p := NewPerson("John", "Smith")
	adapter := NewPersonEmployeeAdapter(p)
	PrintEmployeeName(adapter)
}

```
