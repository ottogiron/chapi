package server

import (
	"fmt"
	"net/http"
)

type Server interface {
	Run(addr string)
	Register(plugin Plugin)
	HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *route
}

type baseServer struct {
	plugins map[string]Plugin
}

func (baseServer *baseServer) Register(plugin Plugin) {

	_, containsKey := baseServer.plugins[plugin.Name()]

	if containsKey {
		panic(fmt.Sprint("Plugin already registered:%s", plugin.Name()))
	}

	baseServer.plugins[plugin.Name()] = plugin
}

func (baseServer *baseServer) registerPlugins(plugins map[string]Plugin, server Server) {

	processedPluginType := new(map[string][]bool)

	for _, plugin := range plugins {
		pluginDendencies := plugin.Dependencies()
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

func pluginDependencyMeet(plugin Plugin, processedPluginType *map[string][]bool) bool {

	return true
}
