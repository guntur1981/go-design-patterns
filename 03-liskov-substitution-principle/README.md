# Description

The Liskov Substitution Principle in Go states that all behaviour of an interface type should be implemented without any problems.

## Example

Consider the following example:

```
package main

import (
	"fmt"
)

type Shape interface {
	SetHeight(h float64)
	SetWidth(w float64)
	GetArea() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (s *Rectangle) SetHeight(h float64) {
	s.height = h
}

func (s *Rectangle) SetWidth(w float64) {
	s.width = w
}

func (s *Rectangle) GetArea() float64 {
	return s.height * s.width
}

func PrintArea(s Shape) {
	fmt.Println("The area of the shape is:", s.GetArea())
}

func main() {
	r := Rectangle{}
	r.SetHeight(2)
	r.SetWidth(3)
	PrintArea(&r)
}
```

In the above example, the Shape interface implemented by Rectangle struct without any problems.

Problems start to arise when we try to implement the Shape interface with a Circle struct, as follows:

```
...

type Circle struct {
	radius float64
}

func (c *Circle) SetHeight(h float64) {
	c.radius = h
}

func (c *Circle) SetWidth(w float64) {
	c.radius = w
}

func (c *Circle) GetArea() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
    ...

	c := Circle{}
	c.SetHeight(4)
	PrintArea(&c)
}


...
```

## The Reason Why This Is Not a Good Practice

The Circle struct has no height or width properties, only a radius. So, the SetWidth() and SetHeight() methods have to be tweaked in such a way so that all behavior of the Shape interface can be implemented.

## A Better Approach

To avoid violation of Liskov Substitution Principle, we must determine how we want to use the type, whether:

- To use it as a generic type, as in the example above, pass it as an argument to the PrintArea(). Or;
- To enforce behavior in implemented struct, like in the example above, SetHeight() and SetWidth().

The risk of violations will be greater if you try to mix the two things above.

So according to the example above, we just want to use the Shape interface as a generic argument to the PrintArea() function.

Better code would be, as follows:

```
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
```
