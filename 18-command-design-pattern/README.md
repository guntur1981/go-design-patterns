# Description

The Command Design Pattern is a behavioral design pattern that turns an instruction into a stand-alone object containing all information about the instruction. This allows you to parameterize methods with different instructions, queue or log instructions, and support undoable operations.

## Example

Consider the following example:

```
package main

import "fmt"

type TextEditor struct {
	content string
}

func (e *TextEditor) Write(text string) {
	e.content += text
	fmt.Println("Current content:", e.content)
}

func (e *TextEditor) Undo() {
	// For simplicity, let's assume undo simply removes the last character
	if len(e.content) > 0 {
		e.content = e.content[:len(e.content)-1]
		fmt.Println("Undoing last modification. Current content:", e.content)
	} else {
		fmt.Println("Nothing to undo.")
	}
}

func main() {
	editor := &TextEditor{}

	editor.Write("Hello ")
	editor.Write("World!")
	editor.Undo()
}
```

The above example demonstrates a sequence of instructions to write and undo a content in `TextEditor`.

## The Reason Why This Is Not a Good Practice

1. **Limited Undo Functionality**: The undo functionality is limited to reverting the last character modification. Because the `TextEditor` does not have information/history of the last instruction that can be undone.
2. **Violates Single Responsibility Principle**: The `TextEditor` handles multiple instructions, text editing and undoing modifications.

## A Better Approach

**First**, let's adhere to Single Responsibility Principle by modifying `TextEditor` to only allow writing or deletion of its content.

```
type TextEditor struct {
	content string
}

func (e *TextEditor) Write(text string) {
	e.content += text
	fmt.Println("Current content:", e.content)
}

func (e *TextEditor) Erase(length int) {
	if len(e.content) >= length {
		e.content = e.content[:len(e.content)-length]
	} else {
		e.content = ""
	}
	fmt.Println("Current content after erase:", e.content)
}
```

**Second**, let's create an interface which represents a set of commands/instructions that can be executed.

```
type Command interface {
	Execute()
	Undo()
}

```

**Third**, implements the interface `Command`:

```
type WriteCommand struct {
	editor *TextEditor
	text   string
}

func (w *WriteCommand) Execute() {
	w.editor.Write(w.text)
}

func (w *WriteCommand) Undo() {
	w.editor.Erase(len(w.text))
}
```

**Fourth**, create a struct that manages a command and executes it. Let's call it an `Invoker`:

```
type Invoker struct {
	history           []Command
	currentHistoryPos int
}

func (i *Invoker) Execute(cmd Command) {
	i.history = append(i.history[:i.currentHistoryPos], cmd)
	i.currentHistoryPos++
	cmd.Execute()
}

func (i *Invoker) Undo() {
	if i.currentHistoryPos > 0 {
		i.currentHistoryPos--
		i.history[i.currentHistoryPos].Undo()
	}
}

func (i *Invoker) Redo() {
	if i.currentHistoryPos < len(i.history) {
		i.history[i.currentHistoryPos].Execute()
		i.currentHistoryPos++
	}
}
```

We also add `Redo()` functionality.

**Finally**, we can implement Command design pattern like this:

```
func main() {
	handler := LowSeverityHandler{}
	handler.
		SetNext(&MediumSeverityHandler{}).
		SetNext(&HighSeverityHandler{})

	// tickets to be handled
	tickets := []SupportTicket{
		{Severity: LowSeverity, Message: "Low severity issue"},
		{Severity: MediumSeverity, Message: "Medium severity issue"},
		{Severity: HighSeverity, Message: "High severity issue"},
	}

	for _, ticket := range tickets {
		handler.Handle(ticket)
	}
}
```
