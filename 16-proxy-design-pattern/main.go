package main

import (
	"fmt"
)

type Image interface {
	Display()
}

type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	image := &RealImage{filename: filename}
	image.loadFromDisk()
	return image
}

func (r *RealImage) loadFromDisk() {
	fmt.Printf("Loading %s from disk\n", r.filename)
}

func (r *RealImage) Display() {
	fmt.Printf("Displaying %s\n", r.filename)
}

// proxy
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

func main() {
	// image := NewRealImage("sample.jpg")
	image := NewProxyImage("sample.jpg")
	image.Display()
}
