package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	BookID    string         `json:"book_id"`
	Judul     string         `json:"judul"`
	Pengarang string         `json:"pengarang"`
	Penerbit  string         `json:"penerbit"`
	Ringkasan string         `json:"ringkasan"`
	Stok      string         `json:"stok"`
	StokMaks  string         `json:"stok_maks"`
	CreatedAt time.Time      `json:"create_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at"`
}

type BorrowBook struct {
	BorrowID   string    `json:"borow_id"`
	BookID     string    `json:"book_id"`
	UserID     string    `json:"user_id"`
	BorrowDate time.Time `json:"borrow_date"`
	DeadLine   time.Time `json:"dead_line"`
	Status     string    `json:"status"`
}

type LendBook struct {
	LendID     string    `json:"lend_id"`
	BookID     string    `json:"book_id"`
	UserID     string    `json:"user_id"`
	BorrowDate time.Time `json:"borrow_date"`
	DeadLine   time.Time `json:"dead_line"`
	LendDate   time.Time `json:"lend_date"`
	Forfeit    int       `json:"forfeit"`
}

type BookUsecaseInterface interface {
	CreateBook(ctx context.Context, b Book) (Book, error)
	UpdateBook(ctx context.Context, b Book, BookID string) (Book, error)
	DeleteBook(ctx context.Context, BookID string) error
	ListBook(ctx context.Context, b Book) (Book, error)
	BorrowedBook(ctx context.Context, bb BorrowBook) (BorrowBook, error)
	ListBorrowBook(ctx context.Context, bb BorrowBook) (BorrowBook, error)
	LendedBook(ctx context.Context, lb LendBook) (LendBook, error)
	ListLendBook(ctx context.Context, lb LendBook) (LendBook, error)
}

type BookRepositoryInterface interface {
	CreateBook(ctx context.Context, b Book) (Book, error)
	UpdateBook(ctx context.Context, b Book, BookID string) (Book, error)
	DeleteBook(ctx context.Context, BookID string) error
	ListBook(ctx context.Context, b Book) (Book, error)
	BorrowBook(ctx context.Context, bb BorrowBook) (BorrowBook, error)
	ListBorrowBook(ctx context.Context, bb BorrowBook) (BorrowBook, error)
	LendedBook(ctx context.Context, lb LendBook) (LendBook, error)
	ListLendBook(ctx context.Context, lb LendBook) (LendBook, error)
}
