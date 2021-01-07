package routes

import (
	"github.com/LoliE1ON/freedom/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetApiRoutes(e *echo.Echo) {

	var api = e.Group("/api")
	api.Use(middleware.BodyLimit("10M"))

	var files = api.Group("/files/avatar")
	{
		files.GET("/download/:id", controller.FilesDownloadAvatar)
		files.POST("/upload", controller.FilesUploadAvatar)
	}

}
