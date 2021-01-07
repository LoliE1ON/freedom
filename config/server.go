package config

import (
	"os"

	"github.com/LoliE1ON/freedom/types"
)

func GetServer() types.ServerConfig {
	return types.ServerConfig{
		Port: os.Getenv("SERVER_PORT"),
	}
}
