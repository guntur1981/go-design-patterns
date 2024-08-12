package main

import "fmt"

type Visitor interface {
	VisitText(*Text)
	VisitImage(*Image)
}

type PrintVisitor struct{}

func (p *PrintVisitor) VisitText(t *Text) {
	fmt.Printf("Print text: %s\n", t.Content)
}

func (p *PrintVisitor) VisitImage(i *Image) {
	fmt.Printf("Print image from path: %s\n", i.Path)
}

type SaveVisitor struct {
	path string
}

func (s *SaveVisitor) VisitText(t *Text) {
	fmt.Printf("Save text: %s to: %s\n", t.Content, s.path)
}

func (s *SaveVisitor) VisitImage(i *Image) {
	fmt.Printf("Save image from path: %s to path: %s\n", i.Path, s.path)
}

type Document interface {
	Accept(v Visitor)
}

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
