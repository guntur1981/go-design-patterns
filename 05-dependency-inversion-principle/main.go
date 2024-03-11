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
