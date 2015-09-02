package server

import (
	"net/http"
	"reflect"
)

type Server interface {
	Run(addr string)
	Register(plugin Plugin)
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *route
}

type baseServer struct {
	plugins []Plugin
}

func (baseServer *baseServer) Register(plugin Plugin) {
	baseServer.plugins = append(baseServer.plugins, plugin)
}

func (baseServer *baseServer) registerPlugins(plugins []Plugin, server Server) {

	processedPluginType := new(map[reflect.Type][]bool)

	for _, plugin := range plugins {
		pluginDendencies := plugin.GetDependencies()
		depLen := len(pluginDendencies)
		if depLen == 0 {
			plugin.Register(server)
		} else {
			meet := pluginDependencyMeet(plugin, processedPluginType)
			if meet {
				plugin.Register(server)
			}
		}

	}
}

func pluginDependencyMeet(plugin Plugin, processedPluginType *map[reflect.Type][]bool) bool {

	return true
}
