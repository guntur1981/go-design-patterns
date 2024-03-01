# Description

The Open-Closed Principle states that a type should be open for extension, but closed for modification.

## Example

Consider the following example:

```
package main

import "fmt"

type Employee struct {
	Department string
	Firstname  string
	Lastname   string
}

type Filter struct{}

func (f Filter) ByDepartment(data []Employee, dept string) []Employee {
	result := []Employee{}
	for _, v := range data {
		if v.Department == dept {
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

	filter := Filter{}

	result := filter.ByDepartment(employees, "Marketing")

	fmt.Printf("%+v\n", result)
}
```

In the above example, the Filter type has only one method for filtering by department.

If we also want to filter by first name, then we have to add another method, such as:

```
...

func (f Filter) ByFirstname(data []Employee, fname string) []Employee {
	result := []Employee{}
	for _, v := range data {
		if v.Firstname == fname {
			result = append(result, v)
		}
	}
	return result
}

...
```

## The Reason Why This Is Not a Good Practice

If we no longer need the filter by first name, instead we want to filter by department and by first name, we have to modify the Filter type, such as:

```
...

// DELETED
// func (f Filter) ByDepartment(data []Employee, dept string) []Employee {
// 	result := []Employee{}
// 	for _, v := range data {
// 		if v.Department == dept {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// func (f Filter) ByFirstname(data []Employee, fname string) []Employee {
// 	result := []Employee{}
// 	for _, v := range data {
// 		if v.Firstname == fname {
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

func (f Filter) ByDepartmentAndFirstname(data []Employee, dept, fname string) []Employee {
	result := []Employee{}
	for _, v := range data {
		if v.Department == dept && v.Firstname == fname {
			result = append(result, v)
		}
	}
	return result
}

...
```

This would violate the Open-Closed Principle where the Filter type should not be opened for modification.

## A Better Approach

A better approach is to change the Filter type to an interface that will be implemented differently based on the requirements, such as:

```
...

type Filter interface {
	Equal(emp *Employee) bool
}

type DepartmentFilter struct {
	Department string
}

func (d DepartmentFilter) Equal(emp *Employee) bool {
	return emp.Department == d.Department
}

...
```

Then, provide a function (method) to use the Filter type depending on needs, such as:

```
...

func FilterEmployees(data []Employee, f Filter) []Employee {
	result := []Employee{}

	for _, v := range data {
		if f.Equal(&v) {
			result = append(result, v)
		}
	}

	return result
}

...
```

Below is the full example source code:

```
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
```
