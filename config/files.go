package config

import (
	"os"
	"strconv"

	"github.com/LoliE1ON/freedom/types"
)

// TODO: нужно добавить возврат ошибки
func GetFiles() types.FilesConfig {

	var filesMaxSize, _ = strconv.Atoi(os.Getenv("FILES_MAX_SIZE"))

	return types.FilesConfig{
		AvatarPath:  os.Getenv("FILES_AVATARS_FOLDER"),
		MaxFileSize: filesMaxSize,
	}
}
