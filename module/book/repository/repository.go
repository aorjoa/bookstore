package repository

import "gorm.io/gorm"

type bookRepository struct {
	Conn *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}
