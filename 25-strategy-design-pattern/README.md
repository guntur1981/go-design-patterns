# Description

The Strategy Design Pattern is a behavioral design pattern that allows the behavior of an object to be selected at runtime.

## Example

Consider the following example:

```
package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	Markdown OutputFormat = iota
	Html
)

type ListProcessor struct {
	outputFormat OutputFormat
	builder      strings.Builder
}

func NewListProcessor() *ListProcessor {
	return &ListProcessor{
		outputFormat: Markdown,
		builder:      strings.Builder{},
	}
}

func (p *ListProcessor) AddItems(items []string) {
	switch p.outputFormat {
	case Markdown:
		for _, item := range items {
			p.builder.WriteString("* " + item + "\n")
		}
	case Html:
		p.builder.WriteString("<ul>\n")
		for _, item := range items {
			p.builder.WriteString("  <li>" + item + "</li>\n")
		}
		p.builder.WriteString("</ul>\n")
	}
}

func (p *ListProcessor) ChangeOutputFormat(fmt OutputFormat) {
	p.outputFormat = fmt
}

func (p *ListProcessor) Reset() {
	p.builder.Reset()
}

func (p *ListProcessor) String() string {
	return p.builder.String()
}

func main() {
	lp := NewListProcessor()
	lp.AddItems([]string{"foo", "bar", "baz"})
	fmt.Println(lp)

	lp.ChangeOutputFormat(Html)
	lp.Reset()
	lp.AddItems([]string{"foo", "bar", "baz"})
	fmt.Println(lp)
}
```

In the above example, we simulate a `ListProcessor` that accepts a list of strings and prints them using a format selected by user.

## The Reason Why This Is Not a Good Practice

1. **Breaks Single Responsibility Principle**: `ListProcessor` is responsible for both handling the format and the logic of each format.
2. **Breaks Open-Closed Principle**: If another format is added, The `AddItems()` method must be modified to handle the new format.

## A Better Approach

**First**, let's change `OutputFormat` constant to an interface:

```
type OutputFormat interface {
	Start(builder *strings.Builder)
	AddItem(builder *strings.Builder, item string)
	End(builder *strings.Builder)
}
```

Here, we enforce common strategies for `Markdown` and `Html` to display list items. These strategies are: `Start()`, `AddItem()`, and `End()` methods.

Of course, for `Markdown` as a concrete object there is no need to implement the `Start()` and `End()` methods so it is sufficient to just leave them empty, as follows:

```
type Markdown struct{}

func (m Markdown) Start(builder *strings.Builder) {
    // do nothing
}

func (m Markdown) AddItem(builder *strings.Builder, item string) {
	builder.WriteString("* " + item + "\n")
}

func (m Markdown) End(builder *strings.Builder) {
    // do nothing
}

```

**Second**, modify `ListProcessor` initialization to accept desired `OutputFormat`.:

```
func NewListProcessor(fmt OutputFormat) *ListProcessor {
	return &ListProcessor{
		outputFormat: fmt,
		builder:      strings.Builder{},
	}
}
```

**Third**, simplify the `AddItems()` method to comply with the Open-Closed Principle, as follows:

```
func (p *ListProcessor) AddItems(items []string) {
	p.outputFormat.Start(&p.builder)
	for _, item := range items {
		p.outputFormat.AddItem(&p.builder, item)
	}
	p.outputFormat.End(&p.builder)
}
```

**Finally**, we can use the Strategy Design Pattern like this:

```
func main() {
	markdown := &Markdown{}
	html := &Html{}

	lp := NewListProcessor(markdown)
	lp.AddItems([]string{"foo", "bar", "baz"})
	fmt.Println(lp)

	lp.ChangeOutputFormat(html)
	lp.Reset()
	lp.AddItems([]string{"foo", "bar", "baz"})
	fmt.Println(lp)
}
```
