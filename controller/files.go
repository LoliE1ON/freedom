package controller

import (
	"net/http"

	"github.com/LoliE1ON/freedom/config"

	"github.com/LoliE1ON/freedom/utils"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func FilesDownloadAvatar(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!!")
}

func FilesUploadAvatar(c echo.Context) error {

	/** Validate fields **/
	id := c.FormValue("id")
	file, err := c.FormFile("file")
	if len(id) == 0 || err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	var newUuid = uuid.Must(uuid.NewRandom()).String()
	var savePath = config.GetConfig().Files.AvatarPath + "/" + newUuid + ".avatar"

	err = utils.UploadFile(file, savePath)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, newUuid)

}
