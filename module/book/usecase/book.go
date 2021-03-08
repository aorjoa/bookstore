package usecase

import (
	"github.com/Aorjoa/bookstore/module/book/converter"
	"github.com/Aorjoa/bookstore/module/book/dto"
	"github.com/Aorjoa/bookstore/module/book/entity"
)

func (b *bookUseCase) GetBookList() ([]dto.BookResponse, error) {
	bl, err := b.bookRepo.GetBookList()
	if err != nil {
		return nil, err
	}
	return converter.ConvertBookEntityListToBookResponseList(bl), err
}

func (b *bookUseCase) GetBookByID(bID uint64) (*dto.BookResponse, error) {
	bl, err := b.bookRepo.GetBookByID(bID)
	if err != nil {
		return nil, err
	}
	return converter.ConvertBookEntityToBookResponse(bl), err
}

func (b *bookUseCase) CreateBook(eb *entity.Book) error {
	return b.bookRepo.CreateBook(eb)
}

func (b *bookUseCase) DeleteBookByID(bID uint64) error {
	return b.bookRepo.DeleteBookByID(bID)
}
