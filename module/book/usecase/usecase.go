package usecase

import (
	"github.com/Aorjoa/bookstore/module/book/repository"
)

type bookUseCase struct {
	bookRepo repository.BookRepository
}

func NewBookUseCase(bookRepo repository.BookRepository) BookUseCase {
	return &bookUseCase{bookRepo}
}
