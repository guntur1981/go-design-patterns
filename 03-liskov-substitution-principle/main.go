package main

import (
	"fmt"
	"math"
)

type Shape interface {
	GetArea() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (r *Rectangle) SetHeight(h float64) {
	r.height = h
}

func (r *Rectangle) SetWidth(w float64) {
	r.width = w
}

func (r *Rectangle) GetArea() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c *Circle) SetRadius(r float64) {
	c.radius = r
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func PrintArea(s Shape) {
	fmt.Println("The area of the shape is:", s.GetArea())
}

func main() {
	r := Rectangle{}
	r.SetHeight(2)
	r.SetWidth(3)
	PrintArea(&r)

	c := Circle{}
	c.SetRadius(4)
	PrintArea(&c)
}
