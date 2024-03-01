package main

import (
	"fmt"
	"os"
	"strings"
)

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

// Not good
func (t Todo) SaveToFile(filename string) error {
	s := strings.Builder{}
	for k, v := range t.items {
		s.WriteString(fmt.Sprintf("%d. %s\n", k+1, v))
	}

	return os.WriteFile(filename, []byte(s.String()), 644)
}

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
