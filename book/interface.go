package book

type BookRepository interface {
	CreateBook(book *Book) error
	GetBookList() ([]Book, error)
	GetBookByID(bID uint64) (*Book, error)
	DeleteBookByID(id uint64) error
}
