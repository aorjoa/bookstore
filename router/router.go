package router

import (
	bookHandler "github.com/Aorjoa/bookstore/module/book/handler"
	bookUseCase "github.com/Aorjoa/bookstore/module/book/usecase"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *fiber.App {
	app := fiber.New()

	// grouping api v1
	v1 := app.Group("/api/v1")

	bg := v1.Group("/books")
	buc := bookUseCase.NewBookUseCase(db)
	bookHandler.NewBookHttpHandler(bg, buc)
	return app
}
