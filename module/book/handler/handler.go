package handler

import (
	"github.com/Aorjoa/bookstore/module/book/usecase"
	"github.com/gofiber/fiber/v2"
)

type HttpBookHandler struct {
	BookUseCase usecase.BookUseCase
}

func NewBookHttpHandler(router fiber.Router, u usecase.BookUseCase) {
	handler := &HttpBookHandler{
		BookUseCase: u,
	}

	router.Get("", handler.GetBookList)
	router.Post("", handler.CreateBook)
}
