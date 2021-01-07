package config

import (
	"os"
)

type Config struct {
	Server ServerType
	Files  FilesType
}

var config Config

func Instantiate() (err error) {

	config = Config{
		Server: ServerType{
			Port:        os.Getenv("SERVER_PORT"),
			MaxBodySize: os.Getenv("FILES_MAX_SIZE"),
		},
		Files: FilesType{
			AvatarPath: os.Getenv("FILES_AVATARS_FOLDER"),
		},
	}

	return
}

func GetConfig() Config {
	return config
}
