package handler

import (
	"github.com/Aorjoa/bookstore/module/book/repository"
	"github.com/Aorjoa/bookstore/module/book/usecase"
	"github.com/gofiber/fiber/v2"
)

type HttpBookHandler struct {
	BookUseCase    usecase.BookUseCase
	BookRepository repository.BookRepository
}

func NewBookHttpHandler(router fiber.Router, u usecase.BookUseCase, br repository.BookRepository) {
	handler := &HttpBookHandler{
		BookUseCase:    u,
		BookRepository: br,
	}

	router.Get("", handler.GetBookList)
	router.Post("", handler.CreateBook)
}
