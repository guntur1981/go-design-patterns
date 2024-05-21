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

type Handler interface {
	SetNext(Handler) Handler
	Handle(SupportTicket)
}

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

type LowSeverityHandler struct {
	BaseHandler
}

func (l *LowSeverityHandler) Handle(ticket SupportTicket) {
	if ticket.Severity == LowSeverity {
		fmt.Printf("Processing low severity ticket: %s\n", ticket.Message)
	} else {
		l.BaseHandler.Handle(ticket)
	}
}

type MediumSeverityHandler struct {
	BaseHandler
}

func (m *MediumSeverityHandler) Handle(ticket SupportTicket) {
	if ticket.Severity == MediumSeverity {
		fmt.Printf("Processing medium severity ticket: %s\n", ticket.Message)
	} else {
		m.BaseHandler.Handle(ticket)
	}
}

type HighSeverityHandler struct {
	BaseHandler
}

func (h *HighSeverityHandler) Handle(ticket SupportTicket) {
	if ticket.Severity == HighSeverity {
		fmt.Printf("Processing high severity ticket: %s\n", ticket.Message)
	} else {
		h.BaseHandler.Handle(ticket)
	}
}

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
