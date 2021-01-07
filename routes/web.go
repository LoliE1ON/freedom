package routes

import (
	"github.com/LoliE1ON/freedom/controller"
	"github.com/labstack/echo/v4"
)

func SetWebRoutes(e *echo.Echo) {

	e.GET("/", controller.HomeIndex)

}
