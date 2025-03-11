package delivery

import "github.com/labstack/echo/v4"

type controller interface {
	GetBooks(c echo.Context) error
	GetBook(c echo.Context) error
	Update(c echo.Context) error
	Create(c echo.Context) error
	Delete(c echo.Context) error
}

func NewRouter(router *echo.Echo, control controller) {
	publicRoutes := router.Group("/api/v1")

	{
		publicRoutes.GET("/", control.GetBooks)
		publicRoutes.GET("/:id", control.GetBooks)
		publicRoutes.POST("/create", control.Create)
		publicRoutes.POST("/update", control.Update)
		publicRoutes.DELETE("/delete", control.Delete)
	}
}
