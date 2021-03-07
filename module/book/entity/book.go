package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name     *string
	ISBN     *string `gorm:"unique"`
	Language *string
	Status   *bool
}
