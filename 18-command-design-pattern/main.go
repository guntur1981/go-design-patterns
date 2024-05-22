package main

import "fmt"

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

type Command interface {
	Execute()
	Undo()
}

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

func main() {
	editor := TextEditor{}

	invoker := Invoker{}
	invoker.Execute(&WriteCommand{editor: &editor, text: "Hello "})
	invoker.Execute(&WriteCommand{editor: &editor, text: "World!"})

	invoker.Undo()
	invoker.Undo()
	invoker.Redo()
	invoker.Redo()
}
