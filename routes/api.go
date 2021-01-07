package routesApi

import (
	"github.com/LoliE1ON/freedom/controller"
	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {

	e.GET("/", controller.HomeIndex)

	files := e.Group("/files")
	{
		files.GET("/get/:id", controller.FilesDownload)
		files.POST("/upload", controller.FilesUpload)
	}

}
