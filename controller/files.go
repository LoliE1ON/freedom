package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func FilesDownload(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}

func FilesUpload(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}
