package main

import "fmt"

type Node interface {
	Name() string
	Ls()
}

// scalar
type File struct {
	name string
}

func (f File) Name() string {
	return f.name
}

func (f File) Ls() {
	fmt.Println("-", f.name)
}

// composite
type Directory struct {
	name     string
	contents []Node
}

func (d Directory) Name() string {
	return d.name
}

func (d Directory) Ls() {
	fmt.Println("Directory:", d.name)
	for _, c := range d.contents {
		c.Ls()
	}
}

func (d *Directory) Add(f Node) {
	d.contents = append(d.contents, f)
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

	dir1.Add(&dir2)

	dir1.Ls()
}
