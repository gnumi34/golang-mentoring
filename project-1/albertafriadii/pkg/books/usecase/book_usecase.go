package usecase

import (
	"context"
	"errors"
	"golang-mentoring/project-1/albertafriadii/pkg/domain"
	"time"
)

type BookUsecase struct {
	UserRepo domain.UserRepositoryInterface
	BookRepo domain.BookRepositoryInterface
}

func NewBookUsecase(r domain.BookRepositoryInterface, u domain.UserRepositoryInterface) domain.BookUsecaseInterface {
	return &BookUsecase{
		UserRepo: u,
		BookRepo: r,
	}
}

func (bc *BookUsecase) ListBook(ctx context.Context) ([]domain.Book, error) {
	books, err := bc.BookRepo.ListBook(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (bc *BookUsecase) CreateBook(ctx context.Context, b *domain.Book) (*domain.Book, error) {
	var err error

	b, err = bc.BookRepo.CreateBook(ctx, b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (bc *BookUsecase) UpdateBook(ctx context.Context, b *domain.Book) error {
	_, err := bc.BookRepo.GetBook(ctx, b.BookID)
	if err != nil {
		return err
	}

	err = bc.BookRepo.UpdateBook(ctx, b)
	if err != nil {
		return err
	}

	return nil
}

func (bc *BookUsecase) DeleteBook(ctx context.Context, BookID uint) error {
	err := bc.BookRepo.DeleteBook(ctx, BookID)
	if err != nil {
		return err
	}

	return nil
}

func (bc *BookUsecase) BorrowBooks(ctx context.Context, bb *domain.BorrowBook) (*domain.BorrowBook, error) {
	book, err := bc.BookRepo.GetBook(ctx, bb.BookID)
	if err != nil {
		return nil, err
	}

	if book.Stock == 0 {
		return nil, errors.New("Out of stock")
	}

	found, err := bc.BookRepo.CheckLend(ctx, bb.BookID, bb.UserID)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.New("Request borrow not found")
	}

	bb.Status = true
	bb.DeadLine = time.Now().AddDate(0, 0, 3)
	res, err := bc.BookRepo.BorrowBooks(ctx, bb)
	if err != nil {
		return nil, err
	}

	book.Stock = book.Stock - 1
	err = bc.BookRepo.UpdateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (bc *BookUsecase) LendBooks(ctx context.Context, lb *domain.LendBook) (*domain.LendBook, error) {
	_, err := bc.BookRepo.GetBook(ctx, lb.BookID)
	if err != nil {
		return nil, err
	}

	lb.Status = false
	res, err := bc.BookRepo.LendBooks(ctx, lb)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (bc *BookUsecase) LendApproval(ctx context.Context, lb *domain.LendBook) (*domain.LendBook, error) {
	_, err := bc.BookRepo.GetBook(ctx, lb.BookID)
	if err != nil {
		return nil, err
	}

	lb.Status = true
	res, err := bc.BookRepo.LendApproval(ctx, lb)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (bc *BookUsecase) ReturnBooks(ctx context.Context, Username string, BookID uint) (*domain.ReturnBook, error) {

	var lateCharge int32

	user, err := bc.UserRepo.FindUserByUsername(ctx, Username)
	if err != nil {
		return nil, err
	}

	book, err := bc.BookRepo.GetBook(ctx, BookID)
	if err != nil {
		return nil, err
	}

	check, err := bc.BookRepo.CheckReturnDate(ctx, user.UserId)
	if err != nil {
		return nil, err
	}

	deadLine := check.DeadLine
	loc, _ := time.LoadLocation("UTC")
	dateNow := time.Now().In(loc)

	if dateNow.After(deadLine) {
		countLate := int(dateNow.Sub(deadLine).Hours() / 24)
		charge := countLate * 5000
		lateCharge = int32(charge)
	}

	status := true
	err = bc.BookRepo.ReturnBooks(ctx, &domain.BorrowBook{
		BookID:     book.BookID,
		UserID:     user.UserId,
		Status:     status,
		ReturnDate: dateNow,
	})
	if err != nil {
		return nil, err
	}

	book.Stock = book.Stock + 1
	err = bc.BookRepo.UpdateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	res := &domain.ReturnBook{
		Username: user.Username,
		Books:    book,
		Charge:   int(lateCharge),
	}

	return res, nil
}

func (bc *BookUsecase) ListBorrowBooks(ctx context.Context, u *domain.GetUserHistory) ([]domain.BorrowBook, error) {
	books, err := bc.BookRepo.ListBorrowBooks(ctx, u.UserID)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (bc *BookUsecase) ListLendBooks(ctx context.Context, u *domain.GetUserHistory) ([]domain.LendBook, error) {
	books, err := bc.BookRepo.ListLendBooks(ctx, u.UserID)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (bc *BookUsecase) ListReturnBooks(ctx context.Context, u *domain.GetUserHistory) ([]domain.BorrowBook, error) {
	books, err := bc.BookRepo.ListReturnBooks(ctx, u.UserID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
