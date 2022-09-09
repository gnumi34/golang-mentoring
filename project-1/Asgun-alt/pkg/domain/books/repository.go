package books

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, domain *BookCollections) (*BookCollections, error)
	UpdateByID(ctx context.Context, domain *BookCollections) error
	DeleteByID(ctx context.Context, domain *BookCollections) error
	GetAll(ctx context.Context) ([]BookCollections, error)

	GetBookByID(ctx context.Context, id uint) (*BookCollections, error)
	GetBookByUserID(ctx context.Context, userID uint, bookID uint) (*BookCollections, error)
	CheckReturnDate(ctx context.Context, userID uint) (*BorrowedBook, error)
	CheckLendRequest(ctx context.Context, bookID, UserID uint) (bool, error)

	BorrowBook(ctx context.Context, domain *BorrowedBook) (*BorrowedBook, error)
	LendApproval(ctx context.Context, domain *LendBook) (*LendBook, error)
	ReturnBook(ctx context.Context, domain *BorrowedBook) error

	ListBorrowedBook(ctx context.Context, userID uint) ([]BorrowedBook, error)
	LendListBook(ctx context.Context, userID uint) ([]LendBook, error)
	ListReturnedBook(ctx context.Context, userID uint) ([]BorrowedBook, error)
}
