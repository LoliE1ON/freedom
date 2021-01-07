package main

import (
	"fmt"
	"log"
	"os"

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

	/** Make config **/
	if err := config.Instantiate(); err != nil {
		log.Fatal("Error parsing config", err)
	}

	/** Create folder for avatars **/
	var avatarDirectory = config.GetConfig().Files.AvatarPath
	if _, err := os.Stat(avatarDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(avatarDirectory, os.ModePerm); err != nil {
			log.Fatal("Failed to create avatars folder", err)
		}
	}

	/** Setup routes **/
	routes.SetApiRoutes(e)
	routes.SetWebRoutes(e)

	/** Start server **/
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.GetConfig().Server.Port)))

}
