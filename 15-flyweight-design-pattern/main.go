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

func main() {
	flyweight := NewFlyweight()
	flyweight.Register("bold", &BoldElement{})
	flyweight.Register("italic", &ItalicElement{})

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
}
