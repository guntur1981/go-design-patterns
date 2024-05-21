# Description

The Chain of Responsibility Design Pattern is a behavioural design pattern that allows an object to pass a request along a chain of potential handlers until the request is handled. Each handler in the chain decides whether to process the request itself or to pass it the next handler in the chain.

## Example

Consider the following example:

```
package main

import "fmt"

type SeverityGrade int

const (
	LowSeverity SeverityGrade = iota
	MediumSeverity
	HighSeverity
)

type SupportTicket struct {
	Severity SeverityGrade
	Message  string
}

func ProcessTicket(ticket SupportTicket) {
	if ticket.Severity <= LowSeverity {
		handleLowSeverity(ticket)
	} else if ticket.Severity == MediumSeverity {
		handleMediumSeverity(ticket)
	} else if ticket.Severity >= HighSeverity {
		handleHighSeverity(ticket)
	} else {
		fmt.Printf("Unknown severity level for ticket: %s\n", ticket.Message)
	}
}

func handleLowSeverity(ticket SupportTicket) {
	fmt.Printf("Processing low severity ticket: %s\n", ticket.Message)
}

func handleMediumSeverity(ticket SupportTicket) {
	fmt.Printf("Processing medium severity ticket: %s\n", ticket.Message)
}

func handleHighSeverity(ticket SupportTicket) {
	fmt.Printf("Processing high severity ticket: %s\n", ticket.Message)
}

func main() {
	// tickets to be handled
	tickets := []SupportTicket{
		{Severity: LowSeverity, Message: "Low severity issue"},
		{Severity: MediumSeverity, Message: "Medium severity issue"},
		{Severity: HighSeverity, Message: "High severity issue"},
	}

	for _, ticket := range tickets {
		ProcessTicket(ticket)
	}
}
```

The above example demonstrates how issues are handled based on their severity.

## The Reason Why This Is Not a Good Practice

1. **Violates Open-Closed Principle**: All the logic the logic for handling different issues is centralized in function `ProcessTicket`. Adding new severity levels or changing the handling logic for existing ones requires modifying this function, which can also lead to a tightly coupled and less maintenable codebase.
2. **Violates Single Responsibility Principle**: The `ProcessTicket` handles multiple responsibilites, deciding which handler to use and processing the ticket.

## A Better Approach

**First**, let's create an interface which can handle all tickets:

```
type Handler interface {
	SetNext(Handler) Handler
	Handle(SupportTicket)
}

```

**Second**, create a struct that implements interface `Handler`:

```
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(h Handler) Handler {
	b.next = h
	return h
}

func (b *BaseHandler) Handle(ticket SupportTicket) {
	if b.next != nil {
		b.next.Handle(ticket)
	}
}
```

**Third**, through embedding, we create all necessary handlers:

```
type LowSeverityHandler struct {
	BaseHandler // embedding
}

type MediumSeverityHandler struct {
	BaseHandler // embedding
}

type HighSeverityHandler struct {
	BaseHandler // embedding
}
```

**Fourth**, extend the behavior of method `Handle()`:

```
func (l *LowSeverityHandler) Handle(ticket SupportTicket) {
	if ticket.Severity == LowSeverity {
		fmt.Printf("Processing low severity ticket: %s\n", ticket.Message)
	} else {
		l.BaseHandler.Handle(ticket)
	}
}

func (m *MediumSeverityHandler) Handle(ticket SupportTicket) {
	if ticket.Severity == MediumSeverity {
		fmt.Printf("Processing medium severity ticket: %s\n", ticket.Message)
	} else {
		m.BaseHandler.Handle(ticket)
	}
}

func (h *HighSeverityHandler) Handle(ticket SupportTicket) {
	if ticket.Severity == HighSeverity {
		fmt.Printf("Processing high severity ticket: %s\n", ticket.Message)
	} else {
		h.BaseHandler.Handle(ticket)
	}
}
```

**Finally**, we can implement Chain of Responsibility design pattern like this:

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
