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
