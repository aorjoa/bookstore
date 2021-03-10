package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type HttpBookHandler struct {
	BookUseCase BookUseCase
}

func NewBookHttpHandler(router fiber.Router, u BookUseCase) {
	handler := &HttpBookHandler{
		BookUseCase: u,
	}

	router.Get("/:bID", handler.GetBookByID)
	router.Get("", handler.GetBookList)
	router.Post("", handler.CreateBook)
	router.Delete("/:bID", handler.DeleteBookByID)
}

func (b *HttpBookHandler) GetBookList(c *fiber.Ctx) error {
	brl, err := b.BookUseCase.GetBookList()
	if err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot get book list")
		return c.Status(http.StatusInternalServerError).SendString("cannot get book list")
	}
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
	br, err := b.BookUseCase.GetBookByID(bID)
	if err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot get book by id")
		return c.Status(http.StatusInternalServerError).SendString("cannot get book by id")
	}
	return c.JSON(br)
}

func (b *HttpBookHandler) CreateBook(c *fiber.Ctx) error {
	var br CreateBookRequest
	if err := c.BodyParser(&br); err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("bad request")
		return c.Status(http.StatusBadRequest).SendString("bad request")
	}

	be := &Book{
		Name:     br.Name,
		ISBN:     br.ISBN,
		Language: br.Language,
		Status:   br.Status,
	}

	if err := b.BookUseCase.CreateBook(be); err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot create book")
		return c.Status(http.StatusInternalServerError).SendString("cannot create book")
	}
	return c.JSON(br)
}

func (b *HttpBookHandler) DeleteBookByID(c *fiber.Ctx) error {
	pbID := c.Params("bID")
	bID, err := strconv.ParseUint(pbID, 0, 64)
	if err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot parse param book id")
		return c.Status(http.StatusInternalServerError).SendString("cannot parse param book id")
	}
	if err := b.BookUseCase.DeleteBookByID(bID); err != nil {
		// TODO: consider to centralize error handler
		log.Err(err).Msg("cannot get book by id")
		return c.Status(http.StatusInternalServerError).SendString("cannot get book by id")
	}
	return c.Status(http.StatusNoContent).JSON(nil)
}
