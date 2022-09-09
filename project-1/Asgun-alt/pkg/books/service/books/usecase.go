package books

import (
	"context"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/books"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/users"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"
	"time"
)

type BooksUseCase struct {
	UserRepo users.UsersRepositoryInterface
	DBRepo   books.Repository
}

func NewBooksUseCase(dbRepo books.Repository, userRepo users.UsersRepositoryInterface) *BooksUseCase {
	return &BooksUseCase{
		UserRepo: userRepo,
		DBRepo:   dbRepo,
	}
}

func (uc *BooksUseCase) GetAllBook(ctx context.Context) ([]books.BookCollections, error) {
	books, err := uc.DBRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (uc *BooksUseCase) AddBook(ctx context.Context, req *books.BookCollections) (*books.BookCollections, error) {
	var err error

	req, err = uc.DBRepo.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (uc *BooksUseCase) UpdateBook(ctx context.Context, req *books.BookCollections) error {
	_, err := uc.DBRepo.GetBookByID(ctx, req.BookID)
	if err != nil {
		return errcode.ErrBookNotFound
	}

	if err = uc.DBRepo.UpdateByID(ctx, req); err != nil {
		return err
	}

	return nil
}

func (uc *BooksUseCase) DeleteBook(ctx context.Context, req *books.BookCollections) error {
	if err := uc.DBRepo.DeleteByID(ctx, req); err != nil {
		return err
	}
	return nil
}

func (uc *BooksUseCase) BorrowBook(ctx context.Context, req *books.BorrowedBook) (*books.BorrowedBook, error) {
	book, err := uc.DBRepo.GetBookByID(ctx, req.BookID)
	if err != nil {
		return nil, errcode.ErrBookNotFound
	}
	if book.Stock == 0 {
		return nil, errcode.ErrStockUnavailable
	}

	found, err := uc.DBRepo.CheckLendRequest(ctx, req.BookID, req.UserID)
	if err != nil {
		return nil, errcode.ErrNotFound
	}
	if !found {
		return nil, errcode.ErrLendRequestNotFound
	}

	req.DueDate = time.Now().AddDate(0, 0, 3)
	response, err := uc.DBRepo.BorrowBook(ctx, req)
	if err != nil {
		return nil, err
	}

	book.Stock -= 1
	if err = uc.DBRepo.UpdateByID(ctx, book); err != nil {
		return nil, errcode.ErrBookNotFound
	}

	return response, nil
}

func (uc *BooksUseCase) LendApproval(ctx context.Context, req *books.LendBook) (*books.LendBook, error) {
	book, err := uc.DBRepo.GetBookByID(ctx, req.BookID)
	if err != nil {
		return nil, errcode.ErrBookNotFound
	}
	if book.Stock == 0 {
		return nil, errcode.ErrStockUnavailable
	}

	// set is_accepted false by default, the status will change to true when admin give approval or it will be set to false if
	// admin didn't give approval and admin will give user a notes to tell the reason behind disproval.
	req.IsAccepted = false
	res, err := uc.DBRepo.LendApproval(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *BooksUseCase) AdminLendApproval(ctx context.Context, req *books.LendBook) (*books.LendBook, error) {
	book, err := uc.DBRepo.GetBookByID(ctx, req.BookID)
	if err != nil {
		return nil, errcode.ErrBookNotFound
	}
	if book.Stock == 0 {
		return nil, errcode.ErrStockUnavailable
	}

	req.IsAccepted = true
	res, err := uc.DBRepo.LendApproval(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (uc *BooksUseCase) ReturnBook(ctx context.Context, username string, bookID uint) (*books.ReturnBookResponse, error) {
	user, err := uc.UserRepo.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	booksResult, err := uc.DBRepo.GetBookByUserID(ctx, user.ID, bookID)
	if err != nil {
		return nil, err
	}

	res, err := uc.DBRepo.CheckReturnDate(ctx, user.ID)
	if err != nil {
		return nil, err

	}

	dueDate := res.DueDate
	loc, _ := time.LoadLocation("UTC")
	currentDate := time.Now().In(loc)
	var lateCharge int32

	if currentDate.After(dueDate) {
		countLate := int(currentDate.Sub(dueDate).Hours() / 24)
		charge := countLate * 5000
		lateCharge = int32(charge)
	}

	hasReturned := true
	err = uc.DBRepo.ReturnBook(ctx, &books.BorrowedBook{
		BookID:     booksResult.BookID,
		UserID:     res.UserID,
		IsReturned: hasReturned,
		ReturnedAt: currentDate,
	})
	if err != nil {
		return nil, err
	}

	booksResult.Stock += 1
	if booksResult.Stock > booksResult.MaxStock {
		return nil, errcode.ErrMaxStockBookLimit
	}
	err = uc.DBRepo.UpdateByID(ctx, booksResult)
	if err != nil {
		return nil, err
	}

	response := &books.ReturnBookResponse{
		Username:   user.Username,
		Books:      booksResult,
		LateCharge: lateCharge,
	}
	return response, nil
}

func (uc *BooksUseCase) ListBorrowedBook(ctx context.Context, req *books.GetUserHistories) ([]books.BorrowedBook, error) {
	books, err := uc.DBRepo.ListBorrowedBook(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (uc *BooksUseCase) LendListBook(ctx context.Context, req *books.GetUserHistories) ([]books.LendBook, error) {
	books, err := uc.DBRepo.LendListBook(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (uc *BooksUseCase) ListReturnedBook(ctx context.Context, req *books.GetUserHistories) ([]books.BorrowedBook, error) {
	books, err := uc.DBRepo.ListReturnedBook(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
