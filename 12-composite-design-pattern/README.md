# Description

The Composite Design Pattern is a mechanism for treating individual (scalar) objects and compositions of objects in a uniform manner.

## Example

Consider the following example:

```
package main

import "fmt"

// scalar
type File struct {
	name string
}

func (f File) Name() string {
	return f.name
}

// composite
type Directory struct {
	name     string
	contents []File
}

func (d Directory) Name() string {
	return d.name
}

func (d Directory) Ls() {
	fmt.Println("Directory:", d.name)
	for _, c := range d.contents {
		fmt.Println(c.Name())
	}
}

func (d *Directory) Add(f *File) {
	d.contents = append(d.contents, *f)
}

func main() {
	file1 := File{name: "file1.txt"}
	file2 := File{name: "file2.txt"}
	dir1 := Directory{name: "dir1"}
	dir1.Add(&file1)
	dir1.Add(&file2)

	img1 := File{name: "image1.jpg"}
	img2 := File{name: "image2.jpg"}
	dir2 := Directory{name: "dir2"}
	dir2.Add(&img1)
	dir2.Add(&img2)

	dir1.Ls()
	dir2.Ls()
}
```

In the above example, `Directory` can only add `File` type to its contents. In reality, we should be able to add sub directories but because the `contents` can only hold `File` type, we cannot do so.

To solve this, we need an interface that can represent both `File` and `Directory` types.

```
type Node interface {
    Name() string
    Ls() string
}
```

Next, add method `Ls()` for type `File` so it can comply with interface `Node`:

```
func (f File) Ls() {
	fmt.Println("-", f.name)
}
```

Next, modify the `contents` slice to allow storing of interface `Node`, and modify method `Add()` accordingly:

```
type Directory struct {
	name     string
	contents []Node // <--
}

...

func (d *Directory) Add(f Node) {
	d.contents = append(d.contents, f)
}
```

Next, modify method `Ls()` so it can recursively print out its contents:

```
func (d Directory) Ls() {
	fmt.Println("Directory:", d.name)
	for _, c := range d.contents {
		c.Ls()
	}
}
```

Finally, we can add `dir2` as subdirectory into `dir1`:

```
func main() {
	file1 := File{name: "file1.txt"}
	file2 := File{name: "file2.txt"}
	dir1 := Directory{name: "dir1"}
	dir1.Add(&file1)
	dir1.Add(&file2)

	img1 := File{name: "image1.jpg"}
	img2 := File{name: "image2.jpg"}
	dir2 := Directory{name: "dir2"}
	dir2.Add(&img1)
	dir2.Add(&img2)

	dir1.Add(&dir2) // <--

	dir1.Ls()
}
```
