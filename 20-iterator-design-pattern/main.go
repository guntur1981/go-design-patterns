package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

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
