package usecase

import "gorm.io/gorm"

type bookUseCase struct {
	db *gorm.DB
}

func NewBookUseCase(db *gorm.DB) BookUseCase {
	return &bookUseCase{db}
}
