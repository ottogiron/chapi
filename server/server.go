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
		alreadyRegisteredMsg := "Plugin already registered:%s"
		panic(fmt.Sprint(alreadyRegisteredMsg, plugin.Name()))
	}

	baseServer.plugins[plugin.Name()] = plugin
}

func (baseServer *baseServer) registerPlugins(plugins map[string]Plugin, server Server) {

	processedPlugins := make(map[string]bool)

	for _, plugin := range plugins {
		dependenciesNames := plugin.Dependencies()
		depLen := len(dependenciesNames)
		if depLen == 0 {
			plugin.Register(server)
		} else {
			meet := pluginDependencyMeet(plugin, processedPlugins)
			if meet {
				plugin.Register(server)
				processedPlugins[plugin.Name()] = true
			} else {

				dependencies := make(map[string]Plugin)
				for _, name := range dependenciesNames {

					if plugins[name] != nil {
						currDep := plugins[plugin.Name()].Name()
						if currDep != plugin.Name() {
							dependencies[name] = plugins[name]
						} else {

						}
					} else {
						unmetDependenciesMsg := "Dependencies for %s are unmet"
						panic(fmt.Sprintf(unmetDependenciesMsg, plugin.Name()))
					}
					baseServer.registerPlugins(dependencies, server)
					plugin.Register(server)
				}
			}
		}

	}
}

func pluginDependencyMeet(plugin Plugin, processedPlugins map[string]bool) bool {

	for _, name := range plugin.Dependencies() {
		if !processedPlugins[name] {
			return false
		}
	}
	return true
}
