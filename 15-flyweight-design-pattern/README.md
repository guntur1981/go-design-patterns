# Description

The Flyweight Design Pattern is a structural design pattern that focuses on minimizing memory usage by sharing as much data as possible with similar objects.

## Example

Consider the following example:

```
package main

import (
	"fmt"
	"strings"
)

type HTMLElement interface {
	Render(content string) string
}

type BoldElement struct{}

func (s *BoldElement) Render(content string) string {
	return fmt.Sprintf("<b>%s</b>", content)
}

type ItalicElement struct{}

func (e *ItalicElement) Render(content string) string {
	return fmt.Sprintf("<i>%s</i>", content)
}

func main() {
	document := []struct {
		tag     string
		content string
	}{
		{"bold", "This is bold text"},
		{"italic", "This is italic text"},
		{"bold", "Another bold text"},
		{"italic", "Another italic text"},
	}

	var renderedDocument strings.Builder
	for _, item := range document {
		var element HTMLElement
		switch item.tag {
		case "bold":
			element = &BoldElement{}
		case "italic":
			element = &ItalicElement{}
		}
		renderedDocument.WriteString(element.Render(item.content))
		renderedDocument.WriteString("\n")
	}

	fmt.Println(renderedDocument.String())
}
```

From the above example, everytime the `renderedDocument` found `bold` or `italic` tag, it will create related object to render the content. This will lead to higher memory consumption due to creating similar objects multiple times.

We might be tempted to modify the codes, as follows:

```
func main() {
	bold := BoldElement{}
	italic := ItalicElement{}

	document := []struct {
		tag     HTMLElement
		content string
	}{
		{&bold, "This is bold text"},
		{&italic, "This is italic text"},
		{&bold, "Another bold text"},
		{&italic, "Another italic text"},
	}

	var renderedDocument strings.Builder
	for _, item := range document {
		renderedDocument.WriteString(item.tag.Render(item.content))
		renderedDocument.WriteString("\n")
	}

	fmt.Println(renderedDocument.String())
}
```

The above code solves the problem of higher memory consumption by avoiding creating similar objects multiple times.

## The Reason Why This Is Not a Good Practice

1. Separation of Concerns (SoC). If the slice `document` sent by other functions, these functions are responsible to create the `HTMLElement` needed for the content, which it should not be.
2. if the SoC is violated, then the problem of higher memory consumption will definitely reappear, because each function will have to create its own `HTMLElement` object (`bold` and `italic`).

## A Better Approach

**First**, let's create a flyweight and its factory:

```
type Flyweight struct {
	elements map[string]HTMLElement
}

func NewFlyweight() *Flyweight {
	return &Flyweight{make(map[string]HTMLElement)}
}

func (f *Flyweight) Register(tag string, element HTMLElement) {
	f.elements[tag] = element
}

func (f *Flyweight) GetElement(tag string) (HTMLElement, error) {
	element, found := f.elements[tag]
	if !found {
		return nil, fmt.Errorf("element %s not found", tag)
	}
	return element, nil
}
```

**Second**, create this factory in client (function `main()`), as well as register its elements:

```
	flyweight := NewFlyweight()
	flyweight.Register("bold", &BoldElement{})
	flyweight.Register("italic", &ItalicElement{})
```

By using this approach, we also adhere to the Open-Closed Principle, where if there are other elements that need to be rendered, we do not need to modify the Flyweight factory.

**Finally**,

```
	document := []struct {
		tag     string
		content string
	}{
		{"bold", "This is bold text"},
		{"italic", "This is italic text"},
		{"bold", "Another bold text"},
		{"italic", "Another italic text"},
	}

	var renderedDocument strings.Builder
	for _, item := range document {
		element, err := flyweight.GetElement(item.tag)
		if err != nil {
			fmt.Println(err)
			continue
		}
		renderedDocument.WriteString(element.Render(item.content))
		renderedDocument.WriteString("\n")
	}

	fmt.Println(renderedDocument.String())
```

we can return `document` to its original form where if this `document` is sent from other functions, these functions do not need to concern about and create the relevant elements (Separation of Concern).
