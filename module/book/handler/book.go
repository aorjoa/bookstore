package handler

import (
	"github.com/Aorjoa/bookstore/module/book/dto"
	"github.com/Aorjoa/bookstore/module/book/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (b *HttpBookHandler) GetBookList(c *fiber.Ctx) error {
	return c.SendString("book list")
}

func (b *HttpBookHandler) CreateBook(c *fiber.Ctx) error {
	var br dto.CreateBookRequest
	if err := c.BodyParser(&br); err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("bad request")
		return c.Status(http.StatusBadRequest).SendString("bad request")
	}

	be := &entity.Book{
		Name:     br.Name,
		ISBN:     br.ISBN,
		Language: br.Language,
		Status:   br.Status,
	}

	if err := b.BookRepository.CreateBook(be); err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot create book")
		return c.Status(http.StatusInternalServerError).SendString("cannot create book")
	}
	return c.JSON(br)
}
