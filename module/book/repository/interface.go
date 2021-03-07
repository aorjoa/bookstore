package repository

import "github.com/Aorjoa/bookstore/module/book/entity"

type BookRepository interface {
	CreateBook(book *entity.Book) error
	GetBookList() ([]entity.Book, error)
}
