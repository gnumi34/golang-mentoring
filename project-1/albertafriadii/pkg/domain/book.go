package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	BookID    uint           `gorm:"primaryKey;auto_increment:true;column:book_id"`
	Title     string         `gorm:"column:title"`
	Author    string         `gorm:"column:author"`
	Publisher string         `gorm:"column:publisher"`
	Summary   string         `gorm:"column:summary"`
	Stock     uint           `gorm:"column:stock"`
	MaxStock  uint           `gorm:"column:max_stock"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type InputCreateBook struct {
	BookID    uint   `json:"book_id"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required,alpha"`
	Publisher string `json:"publisher" validate:"required"`
	Summary   string `json:"summary" validate:"required"`
	Stock     uint   `json:"stock" validate:"required,numeric"`
	MaxStock  uint   `json:"max_stock" validate:"required,numeric"`
}

type InputUpdateBook struct {
	BookID    uint   `json:"book_id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Summary   string `json:"summary"`
	Stock     uint   `json:"stock"`
	MaxStock  uint   `json:"max_stock"`
}

type BorrowBook struct {
	BorrowID   uint           `gorm:"primaryKey;auto_increment:true;column:borrow_id"`
	BookID     uint           `gorm:"foreignKey:BookID;column:book_id"`
	Books      []Book         `gorm:"foreignKey:BookID"`
	UserID     uint           `gorm:"foreignKey:UserID;column:user_id"`
	BorrowDate time.Time      `gorm:"autoCreateTime:true;column:borrow_date"`
	DeadLine   time.Time      `gorm:"column:dead_line"`
	Status     bool           `gorm:"column:status"`
	ReturnDate time.Time      `gorm:"column:return_date"`
	Notes      string         `gorm:"notes"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type InputBorrowBook struct {
	BookID     uint `json:"book_id" validate:"required"`
	UserID     uint `json:"user_id" validate:"required"`
	BorrowDate time.Time
	DeadLine   time.Time
	Status     bool `json:"status"`
	ReturnDate time.Time
	Notes      string `json:"notes"`
}

type LendBook struct {
	LendID      uint           `gorm:"primaryKey;auto_increment:true;column:lend_id"`
	BookID      uint           `gorm:"foreignKey:BookID;column:book_id"`
	UserID      uint           `gorm:"foreignKey:UserID;column:user_id"`
	RequestedAt time.Time      `gorm:"autoCreateTime:true;column:requested_at"`
	Status      bool           `gorm:"column:status"`
	Notes       string         `gorm:"column:notes"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type InputLendBook struct {
	LendID      uint      `json:"lend_id"`
	BookID      uint      `json:"book_id" validate:"required"`
	UserID      uint      `json:"user_id" validate:"required"`
	RequestedAt time.Time `json:"requested_at"`
	Status      bool      `json:"status"`
	Notes       string    `json:"notes"`
}

type ReturnBook struct {
	Username string `json:"username"`
	Books    *Book  `json:"books"`
	Charge   int    `json:"charge"`
}
type InputReturnBook struct {
	Username string `json:"username" validate:"required"`
	BookID   uint   `json:"book_id" validate:"required"`
}

type GetUserHistory struct {
	UserID uint `json:"user_id" validate:"required"`
}

type BookUsecaseInterface interface {
	ListBook(ctx context.Context) ([]Book, error)
	CreateBook(ctx context.Context, b *Book) (*Book, error)
	UpdateBook(ctx context.Context, b *Book) error
	DeleteBook(ctx context.Context, BookID uint) error

	BorrowBooks(ctx context.Context, bb *BorrowBook) (*BorrowBook, error)
	LendBooks(ctx context.Context, lb *LendBook) (*LendBook, error)
	LendApproval(ctx context.Context, lb *LendBook) (*LendBook, error)
	ReturnBooks(ctx context.Context, Username string, BookID uint) (*ReturnBook, error)

	ListBorrowBooks(ctx context.Context, u *GetUserHistory) ([]BorrowBook, error)
	ListLendBooks(ctx context.Context, u *GetUserHistory) ([]LendBook, error)
	ListReturnBooks(ctx context.Context, u *GetUserHistory) ([]BorrowBook, error)
}

type BookRepositoryInterface interface {
	CreateBook(ctx context.Context, b *Book) (*Book, error)
	UpdateBook(ctx context.Context, b *Book) error
	DeleteBook(ctx context.Context, BookID uint) error

	ListBook(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, BookID uint) (*Book, error)
	GetBookByUserID(ctx context.Context, UserID uint, BookID uint) (*Book, error)

	BorrowBooks(ctx context.Context, bb *BorrowBook) (*BorrowBook, error)
	LendBooks(ctx context.Context, lb *LendBook) (*LendBook, error)
	LendApproval(ctx context.Context, lb *LendBook) (*LendBook, error)
	ReturnBooks(ctx context.Context, bb *BorrowBook) error
	CheckReturnDate(ctx context.Context, UserID uint) (*BorrowBook, error)
	CheckLend(ctx context.Context, BookID uint, UserID uint) (bool, error)

	ListBorrowBooks(ctx context.Context, UserId uint) ([]BorrowBook, error)
	ListLendBooks(ctx context.Context, UserId uint) ([]LendBook, error)
	ListReturnBooks(ctx context.Context, UserID uint) ([]BorrowBook, error)
}

func FromBookDomain(b *Book) *Book {
	return &Book{
		BookID:    b.BookID,
		Title:     b.Title,
		Author:    b.Author,
		Publisher: b.Publisher,
		Summary:   b.Summary,
		Stock:     b.Stock,
		MaxStock:  b.MaxStock,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}

func FromBorrowBookDomain(bb *BorrowBook) *BorrowBook {
	return &BorrowBook{
		BorrowID:   bb.BorrowID,
		BookID:     bb.BookID,
		UserID:     bb.UserID,
		BorrowDate: bb.BorrowDate,
		DeadLine:   bb.DeadLine,
		Status:     bb.Status,
		ReturnDate: bb.ReturnDate,
		Notes:      bb.Notes,
	}
}

func FromLendBookDomain(lb *LendBook) *LendBook {
	return &LendBook{
		LendID:      lb.LendID,
		BookID:      lb.BookID,
		UserID:      lb.UserID,
		RequestedAt: lb.RequestedAt,
		Status:      lb.Status,
		Notes:       lb.Notes,
	}
}

func (b *InputCreateBook) ToCreateBookDomain() *Book {
	return &Book{
		BookID:    b.BookID,
		Title:     b.Title,
		Author:    b.Author,
		Publisher: b.Publisher,
		Summary:   b.Summary,
		Stock:     b.Stock,
		MaxStock:  b.MaxStock,
	}
}

func (b *InputUpdateBook) ToUpdateBookDomain() *Book {
	return &Book{
		BookID:    b.BookID,
		Title:     b.Title,
		Author:    b.Author,
		Publisher: b.Publisher,
		Summary:   b.Summary,
		Stock:     b.Stock,
		MaxStock:  b.MaxStock,
	}
}

func (bb *InputBorrowBook) ToBorrowBookDomain() *BorrowBook {
	return &BorrowBook{
		BookID:     bb.BookID,
		UserID:     bb.UserID,
		BorrowDate: bb.BorrowDate,
		DeadLine:   bb.DeadLine,
		Status:     bb.Status,
		ReturnDate: bb.ReturnDate,
		Notes:      bb.Notes,
	}
}

func (lb *InputLendBook) ToLendBookDomain() *LendBook {
	return &LendBook{
		LendID:      lb.LendID,
		BookID:      lb.BookID,
		UserID:      lb.UserID,
		RequestedAt: lb.RequestedAt,
		Status:      lb.Status,
		Notes:       lb.Notes,
	}
}
