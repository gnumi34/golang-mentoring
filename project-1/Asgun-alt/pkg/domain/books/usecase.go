package books

import "context"

type Usecase interface {
	AddBook(ctx context.Context, domain *BookCollections) (*BookCollections, error)
	UpdateBook(ctx context.Context, domain *BookCollections) error
	DeleteBook(ctx context.Context, domain *BookCollections) error
	GetAllBook(ctx context.Context) ([]BookCollections, error)

	BorrowBook(ctx context.Context, domain *BorrowedBook) (*BorrowedBook, error)
	LendApproval(ctx context.Context, domain *LendBook) (*LendBook, error)
	AdminLendApproval(ctx context.Context, domain *LendBook) (*LendBook, error)
	ReturnBook(ctx context.Context, username string, bookID uint) (*ReturnBookResponse, error)

	ListBorrowedBook(ctx context.Context, domain *GetUserHistories) ([]BorrowedBook, error)
	LendListBook(ctx context.Context, domain *GetUserHistories) ([]LendBook, error)
	ListReturnedBook(ctx context.Context, domain *GetUserHistories) ([]BorrowedBook, error)
}
