package book

import (
	"gorm.io/gorm"
)

func ConvertBookEntityListToBookResponseList(bl []Book) []BookResponse {
	var brl []BookResponse
	for _, b := range bl {
		id := b.ID
		br := BookResponse{
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

func ConvertBookEntityToBookResponse(bi *Book) *BookResponse {
	return &BookResponse{
		ID:       &bi.ID,
		Name:     bi.Name,
		ISBN:     bi.ISBN,
		Language: bi.Language,
		Status:   bi.Status,
	}
}

type Book struct {
	gorm.Model
	Name     *string
	ISBN     *string `gorm:"unique"`
	Language *string
	Status   *bool
}
