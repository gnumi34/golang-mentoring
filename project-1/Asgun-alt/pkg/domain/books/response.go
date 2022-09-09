package books

type ReturnBookResponse struct {
	Username   string           `json:"username"`
	Books      *BookCollections `json:"borrowed_books"`
	LateCharge int32            `json:"late_charge"`
}

func FromBooksDomain(domain *BookCollections) *BookCollections {
	return &BookCollections{
		Title:       domain.Title,
		Author:      domain.Author,
		Publisher:   domain.Publisher,
		BookSummary: domain.BookSummary,
		Stock:       domain.Stock,
		MaxStock:    domain.MaxStock,
		Created_At:  domain.Created_At,
		Updated_At:  domain.Updated_At,
		Deleted_At:  domain.Deleted_At,
	}
}

func FromBorrowBookDomain(domain *BorrowedBook) *BorrowedBook {
	return &BorrowedBook{
		BookID:     domain.BookID,
		UserID:     domain.BookID,
		BorrowedAt: domain.BorrowedAt,
		DueDate:    domain.DueDate,
		IsReturned: domain.IsReturned,
		ReturnedAt: domain.ReturnedAt,
		Notes:      domain.Notes,
	}
}

func FromLendAprrovalDomain(domain *LendBook) *LendBook {
	return &LendBook{
		BookID:      domain.BookID,
		UserID:      domain.UserID,
		RequestedAt: domain.RequestedAt,
		IsAccepted:  domain.IsAccepted,
		Notes:       domain.Notes,
	}
}
