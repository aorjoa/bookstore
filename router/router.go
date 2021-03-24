package router

import (
	"github.com/Aorjoa/bookstore/book"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) *fiber.App {
	app := fiber.New()

	// grouping api v1
	v1 := app.Group("/api/v1")

	bg := v1.Group("/books")
	br := book.NewBookRepository(db)
	buc := book.NewBookUseCase(br)
	book.NewBookHttpHandler(bg, buc)
	return app
}
