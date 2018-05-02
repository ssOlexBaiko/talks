// example of DI in Golang

type handler struct {
	storage Storage
}

type Storage interface {
	GetBooks() (storage.Books, error)
	CreateBook(book storage.Book) error
	GetBook(id string) (storage.Book, error)
	RemoveBook(id string) error
	ChangeBook(changedBook storage.Book) (storage.Book, error)
	PriceFilter(filter storage.BookFilter) (storage.Books, error)
	Close() error
}

// Use handler for developing rest-api-server and testing
func NewHandler(storage Storage) *handler {
	return &handler{
		storage: storage,
	}
}
