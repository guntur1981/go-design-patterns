package main

import "fmt"

type Observer interface {
	Notify(componentName, eventName string)
}

type EventLog struct{}

func (e EventLog) Notify(componentName, eventName string) {
	fmt.Printf("%s is %s\n", componentName, eventName)
}

type Button struct {
	name      string
	observers []Observer
}

func (b *Button) Click() {
	// do something on click event
	// ...

	// notify the observer
	b.notifyObservers("clicked")
}

func (b *Button) Subscribe(o Observer) {
	b.observers = append(b.observers, o)
}

func (b *Button) notifyObservers(eventName string) {
	for _, o := range b.observers {
		o.Notify(b.name, eventName)
	}
}

func NewButton(name string) *Button {
	b := &Button{}
	b.name = name
	return b
}

func main() {
	eventLog := &EventLog{}

	submitButton := NewButton("SubmitButton")
	submitButton.Subscribe(eventLog)
	submitButton.Click()
}
