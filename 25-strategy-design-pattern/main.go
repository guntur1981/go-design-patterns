package main

import (
	"fmt"
	"strings"
)

type OutputFormat interface {
	Start(builder *strings.Builder)
	AddItem(builder *strings.Builder, item string)
	End(builder *strings.Builder)
}

type Markdown struct{}

func (m Markdown) Start(builder *strings.Builder) {}

func (m Markdown) AddItem(builder *strings.Builder, item string) {
	builder.WriteString("* " + item + "\n")
}

func (m Markdown) End(builder *strings.Builder) {}

type Html struct{}

func (h Html) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h Html) AddItem(builder *strings.Builder, item string) {
	builder.WriteString("  <li>" + item + "</li>\n")
}

func (h Html) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

type ListProcessor struct {
	outputFormat OutputFormat
	builder      strings.Builder
}

func NewListProcessor(fmt OutputFormat) *ListProcessor {
	return &ListProcessor{
		outputFormat: fmt,
		builder:      strings.Builder{},
	}
}

func (p *ListProcessor) AddItems(items []string) {
	p.outputFormat.Start(&p.builder)
	for _, item := range items {
		p.outputFormat.AddItem(&p.builder, item)
	}
	p.outputFormat.End(&p.builder)
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
