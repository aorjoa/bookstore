package book

type BookUseCase interface {
	GetBookList() ([]BookResponse, error)
	GetBookByID(uint64) (*BookResponse, error)
	CreateBook(*Book) error
	DeleteBookByID(uint64) error
}

type bookUseCase struct {
	bookRepo BookRepository
}

func NewBookUseCase(bookRepo BookRepository) BookUseCase {
	return &bookUseCase{bookRepo}
}

func (b *bookUseCase) GetBookList() ([]BookResponse, error) {
	bl, err := b.bookRepo.GetBookList()
	if err != nil {
		return nil, err
	}
	return ConvertBookEntityListToBookResponseList(bl), err
}

func (b *bookUseCase) GetBookByID(bID uint64) (*BookResponse, error) {
	bl, err := b.bookRepo.GetBookByID(bID)
	if err != nil {
		return nil, err
	}
	return ConvertBookEntityToBookResponse(bl), err
}

func (b *bookUseCase) CreateBook(eb *Book) error {
	return b.bookRepo.CreateBook(eb)
}

func (b *bookUseCase) DeleteBookByID(bID uint64) error {
	return b.bookRepo.DeleteBookByID(bID)
}
