package main

import (
	"github.com/LoliE1ON/freedom/config"
	"github.com/LoliE1ON/freedom/routes"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func main() {

	routesApi.SetRoutes(e)

	var port = config.ServerGetPortForListener()
	e.Logger.Fatal(e.Start(port))
}
