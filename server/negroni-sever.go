package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
)

type negroniServer struct {
	*negroni.Negroni
	router *router
	*baseServer
}

func NewServer() Server {
	n := negroni.Classic()
	router := newRouter()
	baseServer := &baseServer{}
	ns := negroniServer{n, router, baseServer}
	n.UseHandler(router)
	return ns
}

func (server negroniServer) Run(addr string) {
	server.registerPlugins(server.plugins, server)
	server.Negroni.Run(addr)
}

func (server negroniServer) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *route {
	muxRoute := server.router.HandleFunc(path, f)
	route := newRoute(muxRoute)
	return route
}
