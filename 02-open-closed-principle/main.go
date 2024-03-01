package main

import "fmt"

type Employee struct {
	Department string
	Firstname  string
	Lastname   string
}

type Filter interface {
	Equal(emp *Employee) bool
}

type DepartmentFilter struct {
	Department string
}

func (d DepartmentFilter) Equal(emp *Employee) bool {
	return emp.Department == d.Department
}

type FirstnameFilter struct {
	Firstname string
}

func (f FirstnameFilter) Equal(emp *Employee) bool {
	return emp.Firstname == f.Firstname
}

type AndFilter struct {
	First, Second Filter
}

func (a AndFilter) Equal(emp *Employee) bool {
	return a.First.Equal(emp) && a.Second.Equal(emp)
}

func FilterEmployees(data []Employee, f Filter) []Employee {
	result := []Employee{}

	for _, v := range data {
		if f.Equal(&v) {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	employees := []Employee{
		{"Engineering", "Mike", "Whiscard"},
		{"Marketing", "Suzann", "Breeder"},
		{"Marketing", "Shani", "Cranmer"},
		{"Legal", "Nathan", "Hendriks"},
		{"Sales", "Eric", "Stanwood"},
	}

	// filter by department
	deptFilter := DepartmentFilter{"Marketing"}
	result := FilterEmployees(employees, deptFilter)
	fmt.Printf("%+v\n", result)

	// filter by firstname
	fnameFilter := FirstnameFilter{"Suzann"}
	result = FilterEmployees(employees, fnameFilter)
	fmt.Printf("%+v\n", result)

	// filter by department and firstname
	deptFnameFilter := AndFilter{deptFilter, fnameFilter}
	result = FilterEmployees(employees, deptFnameFilter)
	fmt.Printf("%+v\n", result)
}
