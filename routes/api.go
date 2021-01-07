package routes

import (
	"github.com/LoliE1ON/freedom/controller"
	"github.com/labstack/echo/v4"
)

func SetApiRoutes(e *echo.Echo) {

	var api = e.Group("/api")

	var files = api.Group("/files")
	{
		files.GET("/download/:id", controller.FilesDownload)
		files.POST("/upload", controller.FilesUpload)
	}

}
