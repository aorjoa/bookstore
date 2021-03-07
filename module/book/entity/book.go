package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name     *string
	ISBN     *string
	Language *string
	Status   *bool `gorm:"unique"`
}
