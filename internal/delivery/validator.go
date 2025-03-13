package delivery

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BookValidator struct {
	validator *validator.Validate
}

func (cv *BookValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Ошибка валидации: %v", err))
	}
	return nil
}
