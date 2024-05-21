# Description

The Proxy Design Pattern is a structural design pattern that involves using a proxy object to act as an intermediary for another object. This pattern allows you to control access to the original object, perform additional actions before or after accessing the object.

Key differences with Decorator Design Pattern:

- **Proxy Design Pattern:** Used to control access to an object, adding functionalities like access control, or logging without modifying the object itself.
- **Decorator Design Pattern:** Used to add additional behaviour or responsibilites to an object dynamically, providing a flexible alternative to subclassing or extending functionality.

## Example

Consider the following example:

```
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

func main() {
	image := NewRealImage("sample.jpg")
	image.Display()
}
```

From the above example, everytime a `RealImage` is created, it immediately loads the image from disk. This can be inefficient if the image is large or if the image is not needed immediately.

Also, if `RealImage` is created multiple times for the same file, each instance loads the image from disk separately, which is redundant and inefficient.

## A Better Approach

**First**, let's create a proxy object and its factory:

```
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}
```

**Second**, modify method `Display()` to allow lazy initialization:

```
func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}
```

**Finally**, we use the `ProxyImage` for initialization instead of `RealImage`:

```
func main() {
	// image := NewRealImage("sample.jpg")
	image := NewProxyImage("sample.jpg")
	image.Display()
}
```
