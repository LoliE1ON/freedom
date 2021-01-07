package main

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/LoliE1ON/freedom/routes"

	"github.com/LoliE1ON/freedom/config"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func main() {

	/** Load env file **/
	if err := godotenv.Load(".env"); err != nil {
		e.Logger.Fatal("Failed to parse ENV file", err)
	}

	/** Setup routes **/
	routes.SetApiRoutes(e)
	routes.SetWebRoutes(e)

	/** Start server **/
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.GetServer().Port)))

}
