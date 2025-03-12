package delivery

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mizmorr/ingrytech/internal/domain"
)

type Service interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Book, error)
	Create(ctx context.Context, book *domain.Book)
	Update(ctx context.Context, book *domain.Book) error
	GetAll(ctx context.Context) ([]*domain.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type BookController struct {
	svc Service
}

func NewBookController(svc Service) *BookController {
	return &BookController{
		svc: svc,
	}
}

func (bc *BookController) GetBooks(c echo.Context) error {
	books, err := bc.svc.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, books)
}

func (bs *BookController) GetBook(c echo.Context) error {
	rawID := c.Param("id")

	bookID, err := uuid.Parse(rawID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidUUID.Error()})
	}

	book, err := bs.svc.Get(c.Request().Context(), bookID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrInternal.Error()})
	}

	return c.JSON(http.StatusOK, book)
}

func (bs *BookController) Update(c echo.Context) error {
	var (
		bookUpdateReq domain.Book
		err           error
	)

	err = c.Bind(&bookUpdateReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrBadRequestBody.Error()})
	}

	err = bs.svc.Update(c.Request().Context(), &bookUpdateReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrInternal.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated succesfully"})
}

func (bs *BookController) Create(c echo.Context) error {
	var (
		bookCreateReq domain.Book
		err           error
	)

	err = c.Bind(&bookCreateReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrBadRequestBody.Error()})
	}

	bs.svc.Create(c.Request().Context(), &bookCreateReq)

	return c.JSON(http.StatusOK, map[string]string{"status": "created succesfully"})
}

func (bs *BookController) Delete(c echo.Context) error {
	rawID := c.Param("id")

	bookID, err := uuid.Parse(rawID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidUUID.Error()})
	}

	err = bs.svc.Delete(c.Request().Context(), bookID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "deleted succesfully"})
}
