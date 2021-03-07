package handler

import (
	"github.com/Aorjoa/bookstore/module/book/converter"
	"github.com/Aorjoa/bookstore/module/book/dto"
	"github.com/Aorjoa/bookstore/module/book/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (b *HttpBookHandler) GetBookList(c *fiber.Ctx) error {
	bl, err := b.BookRepository.GetBookList()
	if err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot get book list")
		return c.Status(http.StatusInternalServerError).SendString("cannot get book list")
	}

	brl := converter.ConvertBookEntityListToBookResponseList(bl)
	return c.JSON(brl)
}

func (b *HttpBookHandler) GetBookByID(c *fiber.Ctx) error {
	pbID := c.Params("bID")
	bID, err := strconv.ParseUint(pbID, 0, 64)
	if err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot parse param book id")
		return c.Status(http.StatusInternalServerError).SendString("cannot parse param book id")
	}
	bi, err := b.BookRepository.GetBookByID(bID)
	if err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot get book by id")
		return c.Status(http.StatusInternalServerError).SendString("cannot get book by id")
	}

	brl := converter.ConvertBookEntityToBookResponse(bi)
	return c.JSON(brl)
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
