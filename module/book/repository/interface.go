package repository

import "github.com/Aorjoa/bookstore/module/book/entity"

type BookRepository interface {
	CreateBook(book *entity.Book) error
	GetBookList() ([]entity.Book, error)
	GetBookByID(bID uint64) (*entity.Book, error)
	DeleteBookByID(id uint64) error
}
