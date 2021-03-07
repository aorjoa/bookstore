package repository

import (
	"github.com/Aorjoa/bookstore/module/book/entity"
)

func (br *bookRepository) CreateBook(b *entity.Book) error {
	return br.Conn.Create(b).Error
}
