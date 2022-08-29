package books

type AddBookRequest struct {
	Title        string `json:"title" validate:"required, max=50"`
	Author       string `json:"author" validate:"required"`
	Publisher    string `json:"publisher" validate:"required, max=50"`
	BookSummary  string `json:"book_summary" validate:"max=150"`
	BookStock    int    `json:"book_stock" validate:"required"`
	MaxBookStock int    `json:"max_book_stock" validat:"required"`
}

type UpdateBookRequest struct {
	ID           string `json:"id" validate:"required"`
	Title        string `json:"title" validate:"max=50"`
	Author       string `json:"author"`
	Publisher    string `json:"publisher" validate:"max=50"`
	BookSummary  string `json:"book_summary" validate:"max=150"`
	BookStock    int    `json:"book_stock"`
	MaxBookStock int    `json:"max_book_stock"`
}

type ReturnBookRequest struct {
	Username string `json:"username" validate:"required"`
}
