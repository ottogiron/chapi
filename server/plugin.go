package server

import "reflect"

type Plugin interface {
	Register(server Server)
	GetDependencies() []reflect.Type
}

type BasePlugin struct {
}

func (basePlugin *BasePlugin) GetDependencies() []reflect.Type {
	return []reflect.Type{}
}
