package books

import "time"

type AddBookRequest struct {
	Title        string `json:"title" validate:"required,max=50"`
	Author       string `json:"author" validate:"required"`
	Publisher    string `json:"publisher" validate:"required,max=50"`
	BookSummary  string `json:"book_summary" validate:"max=500"`
	BookStock    int    `json:"book_stock" validate:"required,number"`
	MaxBookStock int    `json:"max_book_stock" validat:"required,number"`
}

func (data *AddBookRequest) ToBookDomain() *BookCollections {
	return &BookCollections{
		Title:       data.Title,
		Author:      data.Author,
		Publisher:   data.Publisher,
		BookSummary: data.BookSummary,
		Stock:       data.BookStock,
		MaxStock:    data.BookStock,
	}
}

type UpdateBookRequest struct {
	ID           uint   `json:"id" validate:"required"`
	Title        string `json:"title" validate:"max=50"`
	Author       string `json:"author"`
	Publisher    string `json:"publisher" validate:"max=50"`
	BookSummary  string `json:"book_summary" validate:"max=500"`
	BookStock    int    `json:"book_stock"`
	MaxBookStock int    `json:"max_book_stock"`
}

func (data *UpdateBookRequest) ToBookDomain() *BookCollections {
	return &BookCollections{
		BookID:      data.ID,
		Title:       data.Title,
		Author:      data.Author,
		Publisher:   data.Publisher,
		BookSummary: data.BookSummary,
		Stock:       data.BookStock,
		MaxStock:    data.MaxBookStock,
	}
}

type BorrowBookRequest struct {
	BookID     uint `json:"book_id" validate:"required"`
	UserID     uint `json:"user_id" validate:"required"`
	BorrowedAt time.Time
	DueDate    time.Time
	IsReturned bool `json:"is_returned"`
	ReturnedAt time.Time
	Notes      string `json:"notes" validate:"max=100"`
}

func (data *BorrowBookRequest) ToBookDomain() *BorrowedBook {
	return &BorrowedBook{
		BookID:     data.BookID,
		UserID:     data.UserID,
		BorrowedAt: data.BorrowedAt,
		DueDate:    data.DueDate,
		IsReturned: data.IsReturned,
		ReturnedAt: data.ReturnedAt,
		Notes:      data.Notes,
	}
}

type LendBookRequest struct {
	LendID      uint      `json:"lend_id" validate:"required"`
	BookID      uint      `json:"book_id" validate:"required"`
	UserID      uint      `json:"user_id" validate:"required"`
	RequestedAt time.Time `json:"requested_at"`
	IsAccepted  bool      `json:"is_accepted"`
	Notes       string    `json:"notes" validate:"max=100"`
}

func (data *LendBookRequest) ToBookDomain() *LendBook {
	return &LendBook{
		ID:          data.LendID,
		BookID:      data.BookID,
		UserID:      data.UserID,
		RequestedAt: data.RequestedAt,
		IsAccepted:  data.IsAccepted,
		Notes:       data.Notes,
	}
}

type ReturnBookRequest struct {
	Username string `json:"username" validate:"required"`
	BookID   uint   `json:"book_id" valaidate:"required"`
}

type GetUserHistories struct {
	UserID uint `json:"user_id" validate:"required"`
}
