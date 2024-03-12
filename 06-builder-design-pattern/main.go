package main

import (
	"errors"
	"fmt"
	"strings"
)

// only give access to email behaviour
type Email interface {
	Send() error
}

// hide all properties from user
type email struct {
	from, to, subject, message string
}

func (e email) Send() error {
	// todo: send the email
	return nil
}

type EmailBuilder struct {
	err   error
	email *email
}

func (eb *EmailBuilder) Build() (Email, error) {
	if eb.err != nil {
		return nil, eb.err
	}
	valid := eb.err != nil && len(eb.email.from) > 0 && len(eb.email.to) > 0 && len(eb.email.subject) > 0 && len(eb.email.message) > 0
	if !valid {
		eb.err = errors.New("an email must have from, to, subject and message")
		return nil, eb.err
	}
	return eb.email, nil
}

func (eb *EmailBuilder) From(from string) *EmailBuilder {
	if eb.err != nil {
		return eb
	}
	// validation
	if !strings.Contains(from, "@") {
		eb.err = errors.New("an email address should have '@'")
		return eb
	}

	eb.email.from = from
	return eb
}

func (eb *EmailBuilder) Message(message string) *EmailBuilder {
	if eb.err != nil {
		return eb
	}
	// todo: validation
	eb.email.message = message
	return eb
}

func (eb *EmailBuilder) Reset() *EmailBuilder {
	eb.err = nil
	eb.email = &email{}
	return eb
}

func (eb *EmailBuilder) Subject(subject string) *EmailBuilder {
	if eb.err != nil {
		return eb
	}
	// todo: validation
	eb.email.subject = subject
	return eb
}

func (eb *EmailBuilder) To(to string) *EmailBuilder {
	if eb.err != nil {
		return eb
	}
	// todo: validation
	eb.email.to = to
	return eb
}

func sendMail(from, to, subject, message string) error {
	// validation steps
	if !strings.Contains(from, "@") {
		return errors.New("invalid from email")
	}
	// ... continue the validation steps

	// initialization steps
	e := email{}
	e.from = from
	e.to = to
	e.subject = subject
	e.message = message

	// todo: send the email

	return nil
}

func main() {
	eb := EmailBuilder{}

	// method chaining
	email, err := eb.
		From("foo.com").
		To("bar@test.com").
		Subject("Greeting").
		Message("Hello world!").
		Build()

	if err != nil {
		panic(err)
	}

	email.Send()
	fmt.Println("Email sent successfully!")
}
