package usecase

import (
	"github.com/Aorjoa/bookstore/module/book/dto"
	"github.com/Aorjoa/bookstore/module/book/entity"
)

type BookUseCase interface {
	GetBookList() ([]dto.BookResponse, error)
	GetBookByID(uint64) (*dto.BookResponse, error)
	CreateBook(*entity.Book) error
	DeleteBookByID(uint64) error
}
