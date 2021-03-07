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

func (br *bookRepository) GetBookByID(bID uint64) (*entity.Book, error) {
	var b entity.Book
	err := br.Conn.First(&b, bID).Error
	return &b, err
}

func (br *bookRepository) DeleteBookByID(bID uint64) error {
	err := br.Conn.Delete(&entity.Book{}, bID).Error
	return err
}
