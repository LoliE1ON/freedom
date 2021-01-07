package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server ServerType
	Files  FilesType
}

var config Config

func Instantiate() (err error) {

	filesMaxSize, err := strconv.Atoi(os.Getenv("FILES_MAX_SIZE"))
	if err != nil {
		return
	}

	config = Config{
		Server: ServerType{
			Port: os.Getenv("SERVER_PORT"),
		},
		Files: FilesType{
			AvatarPath:  os.Getenv("FILES_AVATARS_FOLDER"),
			MaxFileSize: filesMaxSize,
		},
	}

	return
}

func GetConfig() Config {
	return config
}
