package books

import (
	"errors"
	"golang-mentoring/project-1/Asgun-alt/cmd/config"
	"golang-mentoring/project-1/Asgun-alt/pkg/common/controller"
	"golang-mentoring/project-1/Asgun-alt/pkg/domain/books"
	"golang-mentoring/project-1/Asgun-alt/pkg/helper/errcode"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BooksHTTPHandler struct {
	common  controller.BaseController
	UseCase books.Usecase
}

func NewBooksHTTPHandler(appGroup *echo.Group, uc books.Usecase) {
	handler := &BooksHTTPHandler{
		UseCase: uc,
	}

	booksGroup := appGroup.Group("/books")
	booksGroup.GET("", handler.GetAllBook)
	booksGroup.POST("/add", handler.AddBook)
	booksGroup.PUT("/update", handler.UpdateBook)
	booksGroup.DELETE("/delete/:id", handler.DeleteBook)

	booksGroup.POST("/borrow_book", handler.BorrowBook)
	booksGroup.POST("/lend_approval", handler.LendApproval)
	booksGroup.POST("/admin_lend_approval", handler.AdminLendApproval)
	booksGroup.PUT("/return_book/:id", handler.ReturnBook)

	booksGroup.GET("/borrow_book_history/:id", handler.GetBorrowedBookHistory)
	booksGroup.GET("/lend_book_history/:id", handler.GetLendBookHistory)
	booksGroup.GET("/returned_book_history/:id", handler.GetReturnedBookHistory)
}

// GetAllBook godoc
// @Summary      get all book
// @Description  get book collections saved in the database.
// @Param 		 id path int true "get all books"
// @Tags         get all books
// @Accept       json
// @Produce      json
// @Success      200  {object}  []books.BookCollections
// @Router       /api/books [get]
func (h *BooksHTTPHandler) GetAllBook(ctx echo.Context) error {
	res, err := h.UseCase.GetAllBook(ctx.Request().Context())
	if err != nil {
		errCode, errMessage := errcode.CheckErrorGetAllBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, res)
}

// AddBook godoc
// @Summary      Add new book
// @Description  Add new book to the database.
// @Param 		 jsonBody body books.AddBookRequest true "Add book body request"
// @Tags         Add book
// @Accept       json
// @Produce      json
// @Success      200  {object}  books.BookCollections
// @Router       /api/books/add [post]
func (h *BooksHTTPHandler) AddBook(ctx echo.Context) error {
	var request books.AddBookRequest
	valid := ctx.Get("validator").(*config.CustomValidator)

	if err := ctx.Bind(&request); err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("failed to process data"))
	}

	validationErr := ctx.Validate(&request)
	if validationErr != nil {
		if val, ok := validationErr.(validator.ValidationErrors); ok {
			translatedErr := val.Translate(valid.Translator)
			return h.common.ErrorValidationResponse(ctx, http.StatusNotFound, translatedErr)
		}
		return h.common.ErrorResponse(ctx, http.StatusNotFound, errors.New("something is wrong with validation"))
	}

	book, err := h.UseCase.AddBook(ctx.Request().Context(), request.ToBookDomain())
	if err != nil {
		errCode, errMessage := errcode.CheckErrorAddBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, book)
}

// UpdateBook godoc
// @Summary      Update book
// @Description  Update book to the database.
// @Param 		 id path int true "update book"
// @Param 		 jsonBody body books.UpdateBookRequest true "Update book body request"
// @Tags         Update book
// @Accept       json
// @Produce      json
// @Success      200  {object}  books.BookCollections
// @Router       /api/books/update [put]
func (h *BooksHTTPHandler) UpdateBook(ctx echo.Context) error {
	var request books.UpdateBookRequest
	valid := ctx.Get("validator").(*config.CustomValidator)

	if err := ctx.Bind(&request); err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("failed to process data"))
	}

	validationErr := ctx.Validate(&request)
	if validationErr != nil {
		if val, ok := validationErr.(validator.ValidationErrors); ok {
			translatedErr := val.Translate(valid.Translator)
			return h.common.ErrorValidationResponse(ctx, http.StatusNotFound, translatedErr)
		}
		return h.common.ErrorResponse(ctx, http.StatusNotFound, errors.New("something is wrong with validation"))
	}

	err := h.UseCase.UpdateBook(ctx.Request().Context(), request.ToBookDomain())
	if err != nil {
		errCode, errMessage := errcode.CheckErrorUpdateBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessOkResponse(ctx)
}

// DeleteBook godoc
// @Summary      Delete book
// @Description  Delete book from the database.
// @Param 		 id path int true "delete book"
// @Tags         Delete book
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /api/books/delete/{id} [delete]
func (h *BooksHTTPHandler) DeleteBook(ctx echo.Context) error {
	var (
		id  int
		err error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user ID"))
	}

	err = h.UseCase.DeleteBook(ctx.Request().Context(), &books.BookCollections{BookID: uint(id)})
	if err != nil {
		errCode, errMessage := errcode.CheckErrorDeleteBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}
	return h.common.SuccessOkResponse(ctx)
}

// BorrowBook godoc
// @Summary      Borrow book
// @Description  Borrow book from the database.
// @Param 		 jsonBody body books.BorrowBookRequest true "borrow book body request"
// @Tags         Borrow book
// @Accept       json
// @Produce      json
// @Success      200  {object}  books.BorrowedBook
// @Router       /api/books/borrow_book [post]
func (h *BooksHTTPHandler) BorrowBook(ctx echo.Context) error {
	var request books.BorrowBookRequest
	valid := ctx.Get("validator").(*config.CustomValidator)

	if err := ctx.Bind(&request); err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("failed to process data"))
	}

	validationErr := ctx.Validate(&request)
	if validationErr != nil {
		if val, ok := validationErr.(validator.ValidationErrors); ok {
			translatedErr := val.Translate(valid.Translator)
			return h.common.ErrorValidationResponse(ctx, http.StatusNotFound, translatedErr)
		}
		return h.common.ErrorResponse(ctx, http.StatusNotFound, errors.New("something is wrong with validation"))
	}

	response, err := h.UseCase.BorrowBook(ctx.Request().Context(), request.ToBookDomain())
	if err != nil {
		errCode, errMessage := errcode.CheckErrorBorrowBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, response)
}

// LendApproval godoc
// @Summary      Lend book approval
// @Description  lend book approval.
// @Param 		 jsonBody body books.LendBookRequest true "lend book body request"
// @Tags         Lend book
// @Accept       json
// @Produce      json
// @Success      200  {object}  books.LendBook
// @Router       /api/books/lend_approval [post]
func (h *BooksHTTPHandler) LendApproval(ctx echo.Context) error {
	var request books.LendBookRequest
	valid := ctx.Get("validator").(*config.CustomValidator)

	if err := ctx.Bind(&request); err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("failed to process data"))
	}

	validationErr := ctx.Validate(&request)
	if validationErr != nil {
		if val, ok := validationErr.(validator.ValidationErrors); ok {
			translatedErr := val.Translate(valid.Translator)
			return h.common.ErrorValidationResponse(ctx, http.StatusNotFound, translatedErr)
		}
		return h.common.ErrorResponse(ctx, http.StatusNotFound, errors.New("something is wrong with validation"))
	}

	response, err := h.UseCase.LendApproval(ctx.Request().Context(), request.ToBookDomain())
	if err != nil {
		errCode, errMessage := errcode.CheckErrorBorrowBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, response)
}

// AdminLendApproval godoc
// @Summary      Admin Lend book approval
// @Description  Admin lend book approval.
// @Param 		 jsonBody body books.LendBookRequest true "lend book body request"
// @Tags         Admin Lend book approval
// @Accept       json
// @Produce      json
// @Success      200  {object}  books.LendBook
// @Router       /api/books/admin_lend_approval [post]
func (h *BooksHTTPHandler) AdminLendApproval(ctx echo.Context) error {
	var request books.LendBookRequest
	if err := ctx.Bind(&request); err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("failed to process data"))
	}

	response, err := h.UseCase.AdminLendApproval(ctx.Request().Context(), request.ToBookDomain())
	if err != nil {
		errCode, errMessage := errcode.CheckErrorBorrowBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, response)
}

// ReturnBook godoc
// @Summary      Return book
// @Description  Return book
// @Param 		 id path int true "return book"
// @Param 		 jsonBody body books.ReturnBookRequest true "return book body request"
// @Tags         Return book
// @Accept       json
// @Produce      json
// @Success      200  {object}  books.ReturnBookResponse
// @Router       /api/books/return_book/{id} [put]
func (h *BooksHTTPHandler) ReturnBook(ctx echo.Context) error {
	var (
		id  int
		err error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user ID"))
	}

	var request books.ReturnBookRequest
	valid := ctx.Get("validator").(*config.CustomValidator)

	if err := ctx.Bind(&request); err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("failed to process data"))
	}

	validationErr := ctx.Validate(&request)
	if validationErr != nil {
		if val, ok := validationErr.(validator.ValidationErrors); ok {
			translatedErr := val.Translate(valid.Translator)
			return h.common.ErrorValidationResponse(ctx, http.StatusNotFound, translatedErr)
		}
		return h.common.ErrorResponse(ctx, http.StatusNotFound, errors.New("something is wrong with validation"))
	}

	request.BookID = uint(id)
	response, err := h.UseCase.ReturnBook(ctx.Request().Context(), request.Username, request.BookID)
	if err != nil {
		errCode, errMessage := errcode.CheckErrorBorrowBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, response)
}

// GetBorrowedBookHistory godoc
// @Summary      Get borrowed book history
// @Description  Get borrowed book history based on user ID
// @Param 		 id path int true "get borrowed book"
// @Tags         Get borrowed book history
// @Accept       json
// @Produce      json
// @Success      200  {object}  []books.BorrowedBook
// @Router       /api/books/borrow_book_history/{id} [get]
func (h *BooksHTTPHandler) GetBorrowedBookHistory(ctx echo.Context) error {
	var (
		id  int
		err error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user ID"))
	}

	res, err := h.UseCase.ListBorrowedBook(ctx.Request().Context(), &books.GetUserHistories{UserID: uint(id)})
	if err != nil {
		errCode, errMessage := errcode.CheckErrorGetAllBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, res)
}

// GetLendBookHistory godoc
// @Summary      Get lend book history
// @Description  Get lend book history based on user ID
// @Param 		 id path int true "get lend book"
// @Tags         Get lend book history
// @Accept       json
// @Produce      json
// @Success      200  {object}  []books.LendBook
// @Router       /api/books/lend_book_history/{id} [get]
func (h *BooksHTTPHandler) GetLendBookHistory(ctx echo.Context) error {
	var (
		id  int
		err error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user ID"))
	}

	res, err := h.UseCase.LendListBook(ctx.Request().Context(), &books.GetUserHistories{UserID: uint(id)})
	if err != nil {
		errCode, errMessage := errcode.CheckErrorGetAllBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, res)
}

// GetReturnedBookHistory godoc
// @Summary      Get returned book history
// @Description  Get returned book history based on user ID
// @Param 		 id path int true "get return book"
// @Tags         Get returned book history
// @Accept       json
// @Produce      json
// @Success      200  {object}  []books.BorrowedBook
// @Router       /api/books/returned_book_history/{id} [get]
func (h *BooksHTTPHandler) GetReturnedBookHistory(ctx echo.Context) error {
	var (
		id  int
		err error
	)

	idStr := ctx.Param("id")
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return h.common.ErrorResponse(ctx, http.StatusBadRequest, errors.New("invalid user ID"))
	}

	res, err := h.UseCase.ListReturnedBook(ctx.Request().Context(), &books.GetUserHistories{UserID: uint(id)})
	if err != nil {
		errCode, errMessage := errcode.CheckErrorGetAllBook(err)
		return h.common.ErrorResponse(ctx, errCode, errMessage)
	}

	return h.common.SuccessDataResponse(ctx, res)
}
