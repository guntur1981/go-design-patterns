# Description

The Visitor Design Pattern is a behavioral design pattern that allows you to add further operations to objects without having to modify them. It is particularly useful when you have a structure of objects that are of different types, and you want to perform various operations on these objects without changing their classes.

## Example

Consider the following example:

```
package main

import "fmt"

type Document interface {
	Print()
}

type Text struct {
	Content string
}

func (t *Text) Print() {
	fmt.Printf("Print text: %s\n", t.Content)
}

type Image struct {
	Path string
}

func (i *Image) Print() {
	fmt.Printf("Print image from path: %s\n", i.Path)
}

func main() {
	documents := []Document{
		&Text{"Hello, World!"},
		&Image{"/images/logo.png"},
	}

	for _, document := range documents {
		document.Print()
	}
}
```

In the above example, we have an abstract `Component` that has a `Print()` method. The Concrete objects `Text` and `Image` are each responsible for implementing their own `Print()` method.

Let's say we want to add a save operation to all components, like so:

```
package main

import "fmt"

type Document interface {
	Print()
	Save(path string)
}

type Text struct {
	Content string
}

func (t *Text) Print() {
	fmt.Printf("Print text: %s\n", t.Content)
}

func (t *Text) Save(path string) {
	fmt.Printf("Save text: %s to: %s\n", t.Content, path)
}

type Image struct {
	Path string
}

func (i *Image) Print() {
	fmt.Printf("Print image from path: %s\n", i.Path)
}

func (i *Image) Save(path string) {
	fmt.Printf("Save image from path: %s to path: %s\n", i.Path, path)
}

func main() {
	documents := []Document{
		&Text{"Hello, World!"},
		&Image{"/images/logo.png"},
	}

	for _, document := range documents {
		document.Print()
		document.Save("/temp/")
	}
}
```

Now all components have the additional responbility to implement the `Save()` method.

## The Reason Why This Is Not a Good Practice

**Breaks Single Responsibility Principle**: All concrete objects of `Component` are responsible for both handling printing and saving.

## A Better Approach

While it is fine to add new functionality, in line with the Open-Closed Principle, it is a good idea to have a separate object that can be responsible for that new functionality.

However, to do this, we must first violate the Open-Closed Principle by enforcing an `Accept()` method to all concrete objects.

**First**, create a framework (interface) `Visitor` which will later implement the desired functions for each concrete object of `Component`.:

```
type Visitor interface {
	VisitText(*Text)
	VisitImage(*Image)
}
```

For the print function, here is the `Visitor` implementation:

```
type PrintVisitor struct{}

func (p *PrintVisitor) VisitText(t *Text) {
	fmt.Printf("Print text: %s\n", t.Content)
}

func (p *PrintVisitor) VisitImage(i *Image) {
	fmt.Printf("Print image from path: %s\n", i.Path)
}
```

Same goes with the save function.

```
type SaveVisitor struct {
	path string
}

func (s *SaveVisitor) VisitText(t *Text) {
	fmt.Printf("Save text: %s to: %s\n", t.Content, s.path)
}

func (s *SaveVisitor) VisitImage(i *Image) {
	fmt.Printf("Save image from path: %s to path: %s\n", i.Path, s.path)
}
```

**Second**, modify `Component` to enforce `Accept()` method:

```
type Document interface {
	Accept(v Visitor)
}
```

Each concrete `Component` object must implement an `Accept()` method which will call the `Visit()` method accordingly.

```
type Text struct {
	Content string
}

func (t *Text) Accept(v Visitor) {
	v.VisitText(t)
}

type Image struct {
	Path string
}

func (i *Image) Accept(v Visitor) {
	v.VisitImage(i)
}
```

**Finally**, we can use the Visitor Design Pattern like this:

```
func main() {
	documents := []Document{
		&Text{"Hello, World!"},
		&Image{"/images/logo.png"},
	}

	printVisitor := PrintVisitor{}
	saveVisitor := SaveVisitor{"/temp/"}

	for _, document := range documents {
		document.Accept(&printVisitor)
		document.Accept(&saveVisitor)
	}
}
```
