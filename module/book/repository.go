package book

import (
	"gorm.io/gorm"
)

type bookRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) CreateBook(b *Book) error {
	return br.Conn.Create(b).Error
}

func (br *bookRepository) GetBookList() ([]Book, error) {
	var bl []Book
	err := br.Conn.Find(&bl).Error
	return bl, err
}

func (br *bookRepository) GetBookByID(bID uint64) (*Book, error) {
	var b Book
	err := br.Conn.First(&b, bID).Error
	return &b, err
}

func (br *bookRepository) DeleteBookByID(bID uint64) error {
	err := br.Conn.Delete(&Book{}, bID).Error
	return err
}
