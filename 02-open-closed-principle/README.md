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

In we also want to filter by first name, then we have to add another method, such as:

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

1. The Todo type is no longer focus on maintaining its items but also to persisting it.

2. The Todo type may be used by several packages. If any package is not satisfied with its persistent functionality, we have to add other functionality or worse, modify the existing one (violating the Open-Close Principle).

## A Better Approach

A better approach is to separate the persistent con into another package or function. For example:

```
func SaveTodo(todo *Todo, filename string) error {
	s := strings.Builder{}
	for i := 0; i < todo.Count(); i++ {
		s.WriteString(fmt.Sprintf("%d. %s\n", i+1, todo.Item(i)))
	}

	return os.WriteFile(filename, []byte(s.String()), 644)
}

func main() {
	todo := Todo{}
	todo.Add("Buy milk.")
	todo.Add("Buy bananas.")
	todo.Add("Go to the cinema.")
	SaveTodo(&todo, "todo.txt")
}
```
