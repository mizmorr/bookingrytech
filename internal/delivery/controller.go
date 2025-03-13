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
	Create(ctx context.Context, book *domain.Book) error
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

// GetBooks godoc
// @Summary Получить список книг
// @Description Возвращает список всех книг
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} domain.Book
// @Failure 400 {object} map[string]string
// @Router /books [get]
func (bc *BookController) GetBooks(c echo.Context) error {
	books, err := bc.svc.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, books)
}

// GetBook godoc
// @Summary Получить книгу по ID
// @Description Возвращает информацию о книге по её идентификатору
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "ID книги"
// @Success 200 {object} domain.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [get]
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

// Update godoc
// @Summary Обновить информацию о книге
// @Description Обновляет данные книги по предоставленной информации
// @Tags books
// @Accept json
// @Produce json
// @Param data body domain.Book true "Данные для обновления книги"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/update [post]
func (bs *BookController) Update(c echo.Context) error {
	var (
		bookUpdateReq domain.Book
		err           error
	)

	err = c.Bind(&bookUpdateReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrBadRequestBody.Error()})
	}

	if err = c.Validate(&bookUpdateReq); err != nil {
		return err
	}

	err = bs.svc.Update(c.Request().Context(), &bookUpdateReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrInternal.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "updated succesfully"})
}

// Create godoc
// @Summary Создать новую книгу
// @Description Создает новую книгу с переданными данными
// @Tags books
// @Accept json
// @Produce json
// @Param data body domain.Book true "Данные для создания книги"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/create [post]
func (bs *BookController) Create(c echo.Context) error {
	var (
		bookCreateReq domain.Book
		err           error
	)

	err = c.Bind(&bookCreateReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrBadRequestBody.Error()})
	}

	if err = c.Validate(&bookCreateReq); err != nil {
		return err
	}

	if err = bs.svc.Create(c.Request().Context(), &bookCreateReq); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrInternal.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "created succesfully"})
}

// Delete godoc
// @Summary Удалить книгу
// @Description Удаляет книгу по её идентификатору
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "ID книги"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [delete]
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
