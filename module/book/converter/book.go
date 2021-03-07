package converter

import (
	"github.com/Aorjoa/bookstore/module/book/dto"
	"github.com/Aorjoa/bookstore/module/book/entity"
)

func ConvertBookEntityListToBookResponseList(bl []entity.Book) []dto.BookResponse {
	var brl []dto.BookResponse
	for _, b := range bl {
		id := b.ID
		br := dto.BookResponse{
			ID:       &id,
			Name:     b.Name,
			ISBN:     b.ISBN,
			Language: b.Language,
			Status:   b.Status,
		}
		brl = append(brl, br)
	}
	return brl
}

func ConvertBookEntityToBookResponse(bi *entity.Book) *dto.BookResponse {
	return &dto.BookResponse{
		ID:       &bi.ID,
		Name:     bi.Name,
		ISBN:     bi.ISBN,
		Language: bi.Language,
		Status:   bi.Status,
	}
}
