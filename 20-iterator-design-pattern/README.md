# Description

The Iterator Design Pattern is a type that facilitates the traversal of various data structures in a controlled manner. Features:

- Keeps a pointer to the current element
- Knows how to move to a different element

## Example

Consider the following example:

```
package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func main() {
	books := []*Book{
		{"The Great Gatsby", "F. Scott Fitzgerald"},
		{"1984", "George Orwell"},
		{"To Kill a Mockingbird", "Harper Lee"},
	}

	// Directly iterate over the collection
	for i := 0; i < len(books); i++ {
		book := books[i]
		fmt.Printf("Book: %s, Author: %s\n", book.Title, book.Author)
	}
}
```

In the above example, we traverse the collection `books` directly.

## The Reason Why This Is Not a Good Practice

**Code Duplication**: Each time you need to traverse the collection, you write similar looping logic.

## A Better Approach

Let's create an iterator for `books`:

```
type BookIterator struct {
	books  []*Book
	curPos int
}

func NewBookIterator(books []*Book) *BookIterator {
	return &BookIterator{books, -1}
}

func (i *BookIterator) Current() *Book {
	return i.books[i.curPos]
}

func (i *BookIterator) Next() bool {
	i.curPos++
	return i.curPos < len(i.books)
}
```

Then, we can use the iterator like this:

```
func main() {
	books := []*Book{
		{"The Great Gatsby", "F. Scott Fitzgerald"},
		{"1984", "George Orwell"},
		{"To Kill a Mockingbird", "Harper Lee"},
	}

	for i := NewBookIterator(books); i.Next(); {
		book := i.Current()
		fmt.Printf("Book: %s, Author: %s\n", book.Title, book.Author)
	}
}
```
