# Description

Single Responsibility Principle in Go states that a type should be made for one specific purpose.

## Example

Consider the following example:

```
package main

type Todo struct {
	items []string
}

func (t *Todo) Add(s string) int {
	t.items = append(t.items, s)
	return len(t.items)
}

func (t Todo) Count() int {
	return len(t.items)
}

func (t *Todo) Delete(index int) string {
	deleted := ""
	if index > -1 && index < len(t.items) {
		deleted = t.items[index]
		copy(t.items[index:], t.items[index+1:])
		t.items = t.items[:len(t.items)-1]
	}
	return deleted
}

func (t Todo) Item(index int) string {
	result := ""
	if index > -1 && index < len(t.items) {
		result = t.items[index]
	}
	return result
}

func main() {
	todo := Todo{}
	todo.Add("Buy milk.")
	todo.Add("Buy bananas.")
	todo.Add("Go to the cinema.")
}
```

In the above example, the Todo type is responsible for managing its items such as: add, delete, get item, count item, etc.

However, it may be tempting to add a functionality, for example to persistent its items.

```
package main

...

func (t Todo) SaveToFile(filename string) error {
	s := strings.Builder{}
	for k, v := range t.items {
		s.WriteString(fmt.Sprintf("%d. %s\n", k+1, v))
	}

	return os.WriteFile(filename, []byte(s.String()), 644)
}

func main() {
	todo := Todo{}
	todo.Add("Buy milk.")
	todo.Add("Buy bananas.")
	todo.Add("Go to the cinema.")
	todo.SaveToFile("todo.txt")
}
```

## The Reason Why This Is Not a Good Practice

1. The Todo type is no longer focus on maintaining its items but also to persisting it.

2. The Todo type may be used by several packages. If any package is not satisfied with its persistent functionality, we have to add other functionality or worse, modify the existing one (violating the Open-Closed Principle).

## A Better Approach

A better approach is to separate the persistent concern into another package or function. For example:

```
package main

...

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
