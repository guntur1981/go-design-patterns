# Description

The Dependency Inversion principle in Go states that high-level modules/packages should not depend on low-level modules/packages except through interfaces.

## Example

Consider the following example:

```
package book

type Book struct {
	Title  string
	Author string
}

var Books []Book
```

```
package main

import (
	"fmt"

	"dip/book"
)

func AddBook(title, author string) *book.Book {
	b := book.Book{title, author}
	book.Books = append(book.Books, b)
	return &b
}

func FindBook(title string) *book.Book {
	for _, v := range book.Books {
		if v.Title == title {
			return &v
		}
	}
	return nil
}

func main() {
	AddBook("Harry Potter", "J.K. Rowling")
	AddBook("A Brief History of Time", "Stephen Hawking")

	b := FindBook("Harry Potter")
	fmt.Printf(b.Author)
}

```

In the example above, the main package (high level) depends on the book package (low level) and interacts directly with the book store within it.

## The Reason Why This Is Not a Good Practice

For small projects where use slice to store books, this may not be an issue. But later, when we decided to use a database to store books, we had to modify a lot of codes, thereby violating the Open-Closed Principle.

## A Better Approach

A better approach is to add an interface between the main package and the book package. So later, if we decide to change the way we store books, we can add another package that can fulfill those needs.

```
package book

type Book struct {
	Title  string
	Author string
}

type BookStorage interface {
	AddBook(title, author string) *Book
	FindBook(title string) *Book
}

type BookSlice struct {
	books []Book
}

func (bs *BookSlice) AddBook(title, author string) *Book {
	bk := Book{title, author}
	bs.books = append(bs.books, bk)
	return &bk
}

func (bs *BookSlice) FindBook(title string) *Book {
	for _, v := range bs.books {
		if v.Title == title {
			return &v
		}
	}
	return nil
}
```

```
package main

import (
	"fmt"

	"dip/book"
)

func addBook1(b book.BookStorage) {
	b.AddBook("Harry Potter", "J.K. Rowling")
}

func addBook2(b book.BookStorage) {
	b.AddBook("A Brief History of Time", "Stephen Hawking")
}

func findBook1(b book.BookStorage) *book.Book {
	return b.FindBook("Harry Potter")
}

func main() {
	books := book.BookSlice{}

	// pretending to use the interface
	addBook1(&books)
	addBook2(&books)
	b := findBook1(&books)

	fmt.Printf(b.Author)
}
```
