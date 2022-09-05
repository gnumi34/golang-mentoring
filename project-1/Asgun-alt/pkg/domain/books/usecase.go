package books

import "context"

type BooksUsecase interface {
	AddBook(ctx context.Context, domain *BookDomain) (*BookDomain, error)
	UpdateBook(ctx context.Context, domain *BookDomain) error
	DeleteBook(ctx context.Context, domain *BookDomain) error
	GetAllBook(ctx context.Context) ([]BookDomain, error)

	BorrowedBook(ctx context.Context, domain *BorrowedBook) (*BorrowedBook, error)
	LendApproval(ctx context.Context, domain *LendBook) (*LendBook, error)
	ReturnBook(ctx context.Context, domain *BorrowedBook) (*ReturnBookResponse, error)

	ListBorrowedBook(ctx context.Context) ([]BorrowedBook, error)
	LendListBook(ctx context.Context) ([]LendBook, error)
	ListReturnedBook(ctx context.Context) ([]BorrowedBook, error)
}
