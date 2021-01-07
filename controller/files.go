package controller

import (
	"net/http"
	"os"

	"github.com/LoliE1ON/freedom/config"
	"github.com/LoliE1ON/freedom/utils"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func FilesDownloadAvatar(c echo.Context) error {

	/** Validate **/
	var requestParam = c.Param("id")
	if _, err := uuid.Parse(requestParam); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	/** Check for file existing **/
	var avatarPath = config.GetConfig().Files.AvatarPath + "/" + requestParam + ".avatar"
	if _, err := os.Stat(avatarPath); os.IsNotExist(err) {
		return c.NoContent(http.StatusNotFound)
	}

	return c.File(avatarPath)
}

func FilesUploadAvatar(c echo.Context) error {

	/** Validate fields **/
	file, err := c.FormFile("file")
	if err != nil {
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
