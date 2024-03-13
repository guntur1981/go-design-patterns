# Description

The Builder Design Pattern allows you to create complex objects **step by step**. In other words, a builder is a separate component that is used to build a complex object.

## Example

Consider the following example:

```
package main

import (
	"errors"
	"fmt"
	"strings"
)

type Email struct {
	from, to, subject, message string
}

func sendMail(from, to, subject, message string) error {
	// validation steps
	if !strings.Contains(from, "@") {
		return errors.New("invalid from email")
	}
	// ... continue the validation steps

	// initialization steps
	e := Email{}
	e.from = from
	e.to = to
	e.subject = subject
	e.message = message

	// todo: send the email

	return nil
}

func main() {
	err := sendMail("foo@test.com", "bar@test.com", "Greeting", "Hello world!")
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent successfully!")
}
```

In the above example, the sendMail function has many steps from validation to creating the Email struct. Then, after all steps have been completed, the actual email sending process is executed.

## The Reason Why This Is Not a Good Practice

1. The example has only four properties that must be set and validated. When dealing with structs with more properties or a growing number of properties, the initialization and validations steps also increase, thus increasing complexity.
2. We want to adhere to Single Responsibility Principle, where the steps for creating and sending email should be separated.

## A Better Approach

Let's create a builder to help initialize the Email struct and validate it as well.

```
type EmailBuilder struct {
	err   error
	email *email
}

func (eb *EmailBuilder) Build() (Email, error) {
	if eb.err != nil {
		return nil, eb.err
	}

	valid := eb.err != nil && len(eb.email.from) > 0 && len(eb.email.to) > 0 &&
		len(eb.email.subject) > 0 && len(eb.email.message) > 0

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
```

Except for the Build method, other methods in the above builder always return the receiver to allow **method chaining**.

To separate the concern, the sendingMail function becomes the email struct behavior:

```
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
```

Notice the difference between `Email` and `email`. To prevent users from creating an email struct without the builder and using the send method, we only expose the interface.

Finally, we use the Builder to create the Email struct and validate it as well:

```
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
```
