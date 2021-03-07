package handler

import (
	"github.com/Aorjoa/bookstore/module/book/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (b *HttpBookHandler) BookList(c *fiber.Ctx) error {
	return c.SendString("book list")
}

func (b *HttpBookHandler) CreateBook(c *fiber.Ctx) error {
	var br dto.CreateBookRequest
	if err := c.BodyParser(&br); err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("bad request")
		return c.Status(http.StatusBadRequest).SendString("bad request")
	}
	return c.JSON(br)
}
