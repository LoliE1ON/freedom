package config

import (
	"github.com/LoliE1ON/freedom/types"
	"strconv"
)

var Server = types.Server{
	Port: 8080,
}

func ServerGetPortForListener() string {
	return ":" + strconv.Itoa(Server.Port)
}
