package configServer

import (
	"github.com/LoliE1ON/freedom/types"
	"strconv"
)

var Server = types.ServerType{
	Port: 8080,
}

func GetPortForListener() string {
	return ":" + strconv.Itoa(Server.Port)
}
