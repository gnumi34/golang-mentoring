package repository

import (
	"context"
	"fmt"
	"golang-mentoring/project-1/albertafriadii/pkg/config"
	"golang-mentoring/project-1/albertafriadii/pkg/domain"

	"gorm.io/gorm"
)

type BookDBRepository struct {
	DB *gorm.DB
}

func NewBookDBRepository(db *gorm.DB) domain.BookRepositoryInterface {
	return &BookDBRepository{
		DB: db,
	}
}

func (r *BookDBRepository) ListBook(ctx context.Context) ([]domain.Book, error) {
	var book []domain.Book

	err := r.DB.WithContext(ctx).Find(&book).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.ListBook: %w", err)
	}

	if book == nil || len(book) == 0 {
		return nil, config.ErrNotFound
	}

	return book, nil
}

func (r *BookDBRepository) GetBook(ctx context.Context, BookID uint) (*domain.Book, error) {
	var book domain.Book

	err := r.DB.WithContext(ctx).First(&book, "book_id = ?", BookID).Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.GetBook: %w", err)
	}

	return &book, nil
}

func (r *BookDBRepository) GetBookByUserID(ctx context.Context, BookID uint, UserID uint) (*domain.Book, error) {
	var book domain.Book

	err := r.DB.Table("borrowbooks").Select("books.book_id, books.title, books.author, books.publisher, books.summary, books.stock, books.max_stock").Joins("INNER JOIN users ON borrowbooks.user_id = users.user_id").Joins("INNER JOIN books ON borrowbooks.book_id = books.book_id").Where("users.user_id = ? AND borrowbooks.book_id = ?", UserID, BookID).Find(&book).Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.GetBookByUserID: %w", err)
	}

	return &book, nil
}

func (r *BookDBRepository) CreateBook(ctx context.Context, b *domain.Book) (*domain.Book, error) {
	book := domain.FromBookDomain(b)

	err := r.DB.WithContext(ctx).Create(&book).Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.Create: %w", err)
	}

	return book, nil
}

func (r *BookDBRepository) UpdateBook(ctx context.Context, b *domain.Book) error {
	err := r.DB.WithContext(ctx).Model(&b).Where("book_id = ?", b.BookID).Updates(&b).Error
	if err != nil {
		return fmt.Errorf("BookDBRepository.UpdateBook: %w", err)
	}

	return nil
}

func (r *BookDBRepository) DeleteBook(ctx context.Context, BookID uint) error {
	var book domain.Book

	err := r.DB.WithContext(ctx).Where("book_id = ?", BookID).Delete(&book).Error
	if err != nil {
		return fmt.Errorf("BookDBRepository.DeleteBook: %w", err)
	}

	return nil
}

func (r *BookDBRepository) BorrowBooks(ctx context.Context, bb *domain.BorrowBook) (*domain.BorrowBook, error) {
	book := domain.FromBorrowBookDomain(bb)

	err := r.DB.WithContext(ctx).Create(&book).Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.BorrowBooks: %w", err)
	}

	return book, nil
}

func (r *BookDBRepository) LendBooks(ctx context.Context, lb *domain.LendBook) (*domain.LendBook, error) {
	book := domain.FromLendBookDomain(lb)
	err := r.DB.WithContext(ctx).Create(&book).Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.LendBooks: %w", err)
	}

	return book, nil
}

func (r *BookDBRepository) LendApproval(ctx context.Context, lb *domain.LendBook) (*domain.LendBook, error) {
	book := domain.FromLendBookDomain(lb)
	err := r.DB.WithContext(ctx).Where("lend_id = ?", lb.LendID).Updates(&book).Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.LendApproval: %w", err)
	}
	return book, nil
}

func (r *BookDBRepository) ReturnBooks(ctx context.Context, bb *domain.BorrowBook) error {
	err := r.DB.Model(&bb).Where("book_id = ?", bb.BookID).Updates(&bb).Error
	if err != nil {
		return fmt.Errorf("BookDBRepository.ReturnBooks: %w", err)
	}

	return nil
}

func (r *BookDBRepository) CheckLend(ctx context.Context, BookID uint, UserID uint) (bool, error) {
	var book domain.LendBook

	err := r.DB.WithContext(ctx).First(&book, "book_id = ? AND user_id = ? AND status = true", BookID, UserID).Error
	if err != nil {
		return false, fmt.Errorf("BookDBRepository.CheckLend: %w", err)
	}

	return true, nil
}

func (r *BookDBRepository) CheckReturnDate(ctx context.Context, UserID uint) (*domain.BorrowBook, error) {
	var book domain.BorrowBook

	err := r.DB.WithContext(ctx).First(&book, "user_id = ?", UserID).Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.CheckReturnDate: %w", err)
	}

	return &book, nil
}

func (r *BookDBRepository) ListBorrowBooks(ctx context.Context, UserID uint) ([]domain.BorrowBook, error) {
	var books []domain.BorrowBook

	err := r.DB.WithContext(ctx).Find(&books, "user_id = ?", UserID).Order("borrow_date ASC").Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.ListBorrowBooks: %w", err)
	}

	if books == nil || len(books) == 0 {
		return nil, config.ErrNotFound
	}

	return books, nil
}

func (r *BookDBRepository) ListLendBooks(ctx context.Context, UserID uint) ([]domain.LendBook, error) {
	var books []domain.LendBook

	err := r.DB.WithContext(ctx).Find(&books, "user_id = ?", UserID).Order("requested_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.ListLendBooks: %w", err)
	}

	if books == nil || len(books) == 0 {
		return nil, config.ErrNotFound
	}

	return books, nil
}

func (r *BookDBRepository) ListReturnBooks(ctx context.Context, UserID uint) ([]domain.BorrowBook, error) {
	var books []domain.BorrowBook

	err := r.DB.WithContext(ctx).Find(&books, "user_id = ?", UserID).Order("return_date ASC").Error
	if err != nil {
		return nil, fmt.Errorf("BookDBRepository.ListReturnBooks: %w", err)
	}

	if books == nil || len(books) == 0 {
		return nil, config.ErrNotFound
	}

	return books, nil
}
