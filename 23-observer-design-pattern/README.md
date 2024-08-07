# Description

The Observer Design Pattern is a behavioral design pattern used to create one-to-many dependency between objects so that when one object (observable) changes, all its dependents (observers) are notified and updated automatically.

## Example

Consider the following example:

```
package main

import "fmt"

type EventLog struct{}

func (e EventLog) Notify(componentName, eventName string) {
	fmt.Printf("%s is %s\n", componentName, eventName)
}

type Button struct {
	name string
	elog *EventLog
}

func (b *Button) Click() {
    // do something on click event
    // ...

    // notify the observer
	if b.elog != nil {
		b.elog.Notify(b.name, "clicked")
	}
}

func NewButton(name string, eventLog *EventLog) *Button {
	b := &Button{}
	b.name = name
	b.elog = eventLog
	return b
}

func main() {
	eventLog := &EventLog{}

	submitButton := NewButton("SubmitButton", eventLog)
	submitButton.Click()
}
```

In the example above, we simulate a Button (observable) triggering a click event and the EventLog (observer) getting notified.

## The Reason Why This Is Not a Good Practice

1. **Breaks Single Responsibility Principle**: The `Button` has two responsibilities: handling click event and notifying the observer.
2. **Breaks Open-Closed Principle**: If another observer is added, The `Button` must notify this new observer by modifiying the `Click()` method.
3. **Tight Coupling**: The observable and observer are tight coupled, leading to less flexible and harder to maintain code.

## A Better Approach

**First**, let's create an interface `Observer` to be used by `Button` (observable):

```
type Observer interface {
	Notify(componentName, eventName string)
}
```

**Second**, add `Subscribe()` method to register an observer:

```
type Button struct {
	name      string
	observers []Observer
}

func (b *Button) Subscribe(o Observer) {
	b.observers = append(b.observers, o)
}

```

**Third**, to adhere Open-Closed Principle, add `notifyObservers()` method to notify all registered observers:

```
func (b *Button) Click() {
	// do something on click event
	// ...

	// notify the observer
	b.notifyObservers("clicked")
}

func (b *Button) notifyObservers(eventName string) {
	for _, o := range b.observers {
		o.Notify(b.name, eventName)
	}
}

```

**Finally**, we can use the Observer Design Pattern like this:

```
func main() {
	eventLog := &EventLog{}

	submitButton := NewButton("SubmitButton")
	submitButton.Subscribe(eventLog)
	submitButton.Click()
}
```
