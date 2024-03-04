# Description

The Liskov Substitution Principle states that the purpose of a type (struct/interface) must be valid in any implementation without breaking the code.

## Example

Consider the following example:

```
package main

import (
	"fmt"
)

type Shape interface {
	SetHeight(h int)
	SetWidth(w int)
	GetArea() int
}

type Rectangle struct {
	height int
	width  int
}

func (s *Rectangle) SetHeight(h int) {
	s.height = h
}

func (s *Rectangle) SetWidth(w int) {
	s.width = w
}

func (s *Rectangle) GetArea() int {
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

In the above example, the Shape interface implemented by Rectangle struct without any problem.

Problems starts to arise when we try to implement the Shape interface with a Circle struct, as follows:

```
...

type Circle struct {
	radius int
}

func (c *Circle) SetHeight(h int) {
	c.radius = h
}

func (c *Circle) SetWidth(w int) {
	c.radius = w
}

func (c *Circle) GetArea() int {
	return int(math.Pi * float64(c.radius) * float64(c.radius))
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

The Circle struct type would violate the Liskov Substitution Principle in several ways:

1. The Circle struct has no height or width properties, only a radius. So, the SetWidth() and SetHeight() methods must be implemented differently to the original purpose of the Shape interface.
2. The area of the Circle struct always returns a fraction number, whereas the GetArea() signature of the Shape interface type return a type int, so a type conversion must be performed and produces less accurate result.

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
