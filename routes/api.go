package routesApi

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoutes(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

}
