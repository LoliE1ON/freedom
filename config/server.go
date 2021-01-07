package configServer

import (
	"github.com/LoliE1ON/freedom/types"
	"strconv"
)

var Server = typesServer.Server{
	Port: 8080,
}

func GetPortForListener() string {
	return ":" + strconv.Itoa(Server.Port)
}
