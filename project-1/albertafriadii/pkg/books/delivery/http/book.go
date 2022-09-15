package http

import (
	"errors"
	"golang-mentoring/project-1/albertafriadii/pkg/config"
	"golang-mentoring/project-1/albertafriadii/pkg/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	usecase domain.BookUsecaseInterface
}

func NewBookController(bu domain.BookUsecaseInterface) *BookController {
	return &BookController{
		usecase: bu,
	}
}

func (c *BookController) ListBook(ctx echo.Context) error {
	book, err := c.usecase.ListBook(ctx.Request().Context())
	if err != nil {
		if errors.Is(err, config.ErrNotFound) {
			return domain.ErrResponse(ctx, http.StatusNotFound, config.ErrNotFound)
		}

		return domain.ErrResponse(ctx, http.StatusInternalServerError, config.ErrInternalServerError)
	}

	return domain.SuccessDataResponse(ctx, book)
}

func (c *BookController) CreateBook(ctx echo.Context) error {
	var b domain.InputCreateBook
	var err error

	bindErr := ctx.Bind(&b)
	if bindErr != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("data bind error"))
	}

	validationErr := config.Validator(&b)
	if validationErr != nil {
		return domain.ErrValidResponse(ctx, http.StatusBadRequest, validationErr)
	}

	if !config.ValidationBook(b.Author) {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Pengarang Must Uppercase"))
	}

	_, err = c.usecase.CreateBook(ctx.Request().Context(), b.ToCreateBookDomain())
	if err != nil {
		errStatus, errMessage := config.ErrInputBookCheck(err)
		return domain.ErrResponse(ctx, errStatus, errMessage)
	}
	return domain.OkResponse(ctx)
}

func (c *BookController) UpdateBook(ctx echo.Context) error {
	var b domain.InputUpdateBook
	var id int
	var err error

	BookID := ctx.Param("book_id")
	id, err = strconv.Atoi(BookID)
	if err != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Invalid ID"))
	}

	bindErr := ctx.Bind(&b)
	if bindErr != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("data bind error"))
	}

	validationErr := config.Validator(b)
	if validationErr != nil {
		return domain.ErrValidResponse(ctx, http.StatusBadRequest, validationErr)
	}

	if uint(id) != b.BookID {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Invalid ID"))
	}

	if b.Author != "" {
		if !config.ValidationBook(b.Author) {
			return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Must Uppercase"))
		}
	} else {
		b.Author = ""
	}

	err = c.usecase.UpdateBook(ctx.Request().Context(), b.ToUpdateBookDomain())
	if err != nil {
		errStatus, errMessage := config.ErrInputBookCheck(err)
		return domain.ErrResponse(ctx, errStatus, errMessage)
	}
	return domain.OkResponse(ctx)
}

func (c *BookController) DeleteBook(ctx echo.Context) error {
	var id int
	var err error

	BookID := ctx.Param("book_id")
	id, err = strconv.Atoi(BookID)
	if err != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Invalid ID"))
	}

	err = c.usecase.DeleteBook(ctx.Request().Context(), uint(id))
	if err != nil {
		errStatus, errMessage := config.ErrDeleteCheck(err)
		return domain.ErrResponse(ctx, errStatus, errMessage)
	}
	return domain.OkResponse(ctx)
}

func (c *BookController) BorrowBooks(ctx echo.Context) error {
	var b domain.InputBorrowBook

	bindErr := ctx.Bind(&b)
	if bindErr != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("data bind error"))
	}

	validationErr := config.Validator(b)
	if validationErr != nil {
		return domain.ErrValidResponse(ctx, http.StatusBadRequest, validationErr)
	}

	res, err := c.usecase.BorrowBooks(ctx.Request().Context(), b.ToBorrowBookDomain())
	if err != nil {
		errStatus, errMessage := config.ErrBorrowBookCheck(err)
		return domain.ErrResponse(ctx, errStatus, errMessage)
	}

	return domain.SuccessDataResponse(ctx, res)
}

func (c *BookController) LendBooks(ctx echo.Context) error {
	var b domain.InputLendBook
	bindErr := ctx.Bind(&b)
	if bindErr != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("data bind error"))
	}

	validationErr := config.Validator(b)
	if validationErr != nil {
		return domain.ErrValidResponse(ctx, http.StatusBadRequest, validationErr)
	}

	res, err := c.usecase.LendBooks(ctx.Request().Context(), b.ToLendBookDomain())
	if err != nil {
		errStatus, errMessage := config.ErrBorrowBookCheck(err)
		return domain.ErrResponse(ctx, errStatus, errMessage)
	}

	return domain.SuccessDataResponse(ctx, res)
}

func (c *BookController) LendApproval(ctx echo.Context) error {
	var b domain.InputLendBook
	bindErr := ctx.Bind(&b)
	if bindErr != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("data bind error"))
	}

	res, err := c.usecase.LendApproval(ctx.Request().Context(), b.ToLendBookDomain())
	if err != nil {
		errStatus, errMessage := config.ErrBorrowBookCheck(err)
		return domain.ErrResponse(ctx, errStatus, errMessage)
	}

	return domain.SuccessDataResponse(ctx, res)
}

func (c *BookController) ReturnBooks(ctx echo.Context) error {
	var b domain.InputReturnBook
	var id int
	var err error

	bookID := ctx.Param("book_id")
	id, err = strconv.Atoi(bookID)
	if err != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Invalid ID"))
	}

	bindErr := ctx.Bind(&b)
	if bindErr != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("data bind error"))
	}

	validationErr := config.Validator(b)
	if validationErr != nil {
		return domain.ErrValidResponse(ctx, http.StatusBadRequest, validationErr)
	}

	b.BookID = uint(id)
	res, err := c.usecase.ReturnBooks(ctx.Request().Context(), b.Username, b.BookID)
	if err != nil {
		errStatus, errMessage := config.ErrBorrowBookCheck(err)
		return domain.ErrResponse(ctx, errStatus, errMessage)
	}

	return domain.SuccessDataResponse(ctx, res)
}

func (c *BookController) BorrowBookHistory(ctx echo.Context) error {
	var id int
	var err error

	userID := ctx.Param("user_id")
	id, err = strconv.Atoi(userID)
	if err != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Invalid ID"))
	}

	res, err := c.usecase.ListBorrowBooks(ctx.Request().Context(), &domain.GetUserHistory{
		UserID: uint(id),
	})
	if err != nil {
		if errors.Is(err, config.ErrNotFound) {
			return domain.ErrResponse(ctx, http.StatusNotFound, config.ErrNotFound)
		}

		return domain.ErrResponse(ctx, http.StatusInternalServerError, config.ErrInternalServerError)
	}

	return domain.SuccessDataResponse(ctx, res)
}

func (c *BookController) LendBookHistory(ctx echo.Context) error {
	var id int
	var err error

	userID := ctx.Param("user_id")
	id, err = strconv.Atoi(userID)
	if err != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Invalid ID"))
	}

	res, err := c.usecase.ListLendBooks(ctx.Request().Context(), &domain.GetUserHistory{
		UserID: uint(id),
	})
	if err != nil {
		if errors.Is(err, config.ErrNotFound) {
			return domain.ErrResponse(ctx, http.StatusNotFound, config.ErrNotFound)
		}

		return domain.ErrResponse(ctx, http.StatusInternalServerError, config.ErrInternalServerError)
	}

	return domain.SuccessDataResponse(ctx, res)
}

func (c *BookController) ReturnBookHistory(ctx echo.Context) error {
	var id int
	var err error

	userID := ctx.Param("user_id")
	id, err = strconv.Atoi(userID)
	if err != nil {
		return domain.ErrResponse(ctx, http.StatusBadRequest, errors.New("Invalid ID"))
	}

	res, err := c.usecase.ListReturnBooks(ctx.Request().Context(), &domain.GetUserHistory{
		UserID: uint(id),
	})
	if err != nil {
		if errors.Is(err, config.ErrNotFound) {
			return domain.ErrResponse(ctx, http.StatusNotFound, config.ErrNotFound)
		}

		return domain.ErrResponse(ctx, http.StatusInternalServerError, config.ErrInternalServerError)
	}

	return domain.SuccessDataResponse(ctx, res)
}
