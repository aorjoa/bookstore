package repository

import (
	"github.com/Aorjoa/bookstore/module/book/entity"
)

func (br *bookRepository) CreateBook(b *entity.Book) error {
	return br.Conn.Create(b).Error
}

func (br *bookRepository) GetBookList() ([]entity.Book, error) {
	var bl []entity.Book
	err := br.Conn.Find(&bl).Error
	return bl, err
}
