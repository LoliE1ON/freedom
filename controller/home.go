package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomeIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
