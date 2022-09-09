package db

import (
	"context"
	"fmt"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/books"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"

	"gorm.io/gorm"
)

type BooksDBRepository struct {
	DB *gorm.DB
}

func NewBooksDBRepository(gormDB *gorm.DB) *BooksDBRepository {
	return &BooksDBRepository{DB: gormDB}
}

func (r *BooksDBRepository) CheckLendRequest(ctx context.Context, bookID, UserID uint) (bool, error) {
	var response books.LendBook

	if err := r.DB.WithContext(ctx).First(&response, "book_id = ? AND user_id = ? AND is_accepted = true", bookID, UserID).Error; err != nil {
		return false, fmt.Errorf("BooksDBRepository.GetBookByID: %w", err)
	}

	return true, nil
}

func (r *BooksDBRepository) CheckReturnDate(ctx context.Context, userID uint) (*books.BorrowedBook, error) {
	var response books.BorrowedBook

	if err := r.DB.WithContext(ctx).First(&response, "user_id = ?", userID).Error; err != nil {
		return nil, fmt.Errorf("BooksDBRepository.GetBookByID: %w", err)
	}

	return &response, nil
}

func (r *BooksDBRepository) GetBookByUserID(ctx context.Context, userID uint, bookID uint) (*books.BookCollections, error) {
	var (
		response books.BookCollections
	)

	err := r.DB.Table("borrowed_books").
		Select("*").
		Joins("INNER JOIN users ON borrowed_books.user_id = users.id").
		Joins("INNER JOIN book_collections ON borrowed_books.book_id = book_collections.book_id").
		Where("users.id = ? AND borrowed_books.book_id = ?", userID, bookID).
		Unscoped().
		Find(&response)
	if err.Error != nil {
		fmt.Println("error: ", err)
		return nil, errcode.ErrRecordNotFound
	}

	return &response, nil
}

func (r *BooksDBRepository) GetBookByID(ctx context.Context, id uint) (*books.BookCollections, error) {
	var response books.BookCollections

	if err := r.DB.WithContext(ctx).First(&response, "book_id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("BooksDBRepository.GetBookByID: %w", err)
	}

	return &response, nil
}

func (r *BooksDBRepository) GetAll(ctx context.Context) ([]books.BookCollections, error) {
	var res []books.BookCollections

	err := r.DB.WithContext(ctx).Find(&res).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("BooksDBRepository.GetAll: %w", err)
	}

	if res == nil {
		return nil, errcode.ErrRecordNotFound
	}
	return res, nil
}

func (r *BooksDBRepository) Create(ctx context.Context, req *books.BookCollections) (*books.BookCollections, error) {
	res := books.FromBooksDomain(req)

	err := r.DB.WithContext(ctx).Create(&res).Error
	if err != nil {
		return nil, fmt.Errorf("BooksDBRepository.Create: %w", err)
	}
	return res, nil
}

func (r *BooksDBRepository) UpdateByID(ctx context.Context, req *books.BookCollections) error {
	if err := r.DB.Model(&req).Where("book_id = ?", req.BookID).Updates(&req).Error; err != nil {
		return fmt.Errorf("BooksDBRepository.UpdateByID: %w", err)
	}
	return nil
}

func (r *BooksDBRepository) DeleteByID(ctx context.Context, req *books.BookCollections) error {
	if err := r.DB.WithContext(ctx).Where("book_id = ?", req.BookID).Delete(req).Error; err != nil {
		return fmt.Errorf("BooksDBRepository.DeleteByID: %w", err)
	}
	return nil
}

func (r *BooksDBRepository) BorrowBook(ctx context.Context, req *books.BorrowedBook) (*books.BorrowedBook, error) {
	res := books.FromBorrowBookDomain(req)

	err := r.DB.WithContext(ctx).Create(&res).Error
	if err != nil {
		return nil, fmt.Errorf("BooksDBRepository.BorrowBook: %w", err)
	}
	return res, nil
}

func (r *BooksDBRepository) LendApproval(ctx context.Context, req *books.LendBook) (*books.LendBook, error) {
	res := books.FromLendAprrovalDomain(req)
	err := r.DB.WithContext(ctx).Create(&res).Error
	if err != nil {
		return nil, fmt.Errorf("BooksDBRepository.LendApproval: %w", err)
	}
	return res, nil
}

func (r *BooksDBRepository) ReturnBook(ctx context.Context, req *books.BorrowedBook) error {
	if err := r.DB.Model(&req).Where("book_id = ?", req.BookID).Updates(&req).Error; err != nil {
		return fmt.Errorf("BooksDBRepository.ReturnBook: %w", err)
	}
	return nil
}

func (r *BooksDBRepository) ListBorrowedBook(ctx context.Context, userID uint) ([]books.BorrowedBook, error) {
	var res []books.BorrowedBook

	err := r.DB.WithContext(ctx).Find(&res, "user_id = ?", userID).Order("borrowed_at ASC").Error
	if err != nil {
		return nil, fmt.Errorf("BooksDBRepository.ListBorrowedBook: %w", err)
	}

	if res == nil {
		return nil, errcode.ErrRecordNotFound
	}
	return res, nil
}

func (r *BooksDBRepository) LendListBook(ctx context.Context, userID uint) ([]books.LendBook, error) {
	var res []books.LendBook

	err := r.DB.WithContext(ctx).Find(&res, "user_id = ?", userID).Order("created_at DESC").Error
	if err != nil {
		return nil, fmt.Errorf("BooksDBRepository.LendListBook: %w", err)
	}

	if res == nil {
		return nil, errcode.ErrRecordNotFound
	}
	return res, nil
}

func (r *BooksDBRepository) ListReturnedBook(ctx context.Context, userID uint) ([]books.BorrowedBook, error) {
	var res []books.BorrowedBook

	err := r.DB.WithContext(ctx).Find(&res, "user_id = ?", userID).Order("requested_at ASC").Error
	if err != nil {
		return nil, fmt.Errorf("BooksDBRepository.ListReturnedBook: %w", err)
	}

	if res == nil {
		return nil, errcode.ErrRecordNotFound
	}
	return res, nil
}
