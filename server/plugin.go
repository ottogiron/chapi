package server

import (
	"fmt"

	"github.com/satori/go.uuid"
)

type Plugin interface {
	Register(server Server)
	Dependencies() []string
	Name() string
}

type BasePlugin struct {
}

func (basePlugin *BasePlugin) Dependencies() []string {
	fmt.Println("Calling this from base")
	return []string{}
}

func (basePlugin *BasePlugin) Name() string {
	return uuid.NewV4().String()
}
