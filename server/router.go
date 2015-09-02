package server

import "github.com/gorilla/mux"

type router struct {
	*mux.Router
}

func newRouter() *router {
	negroniRouter := mux.NewRouter()
	return &router{negroniRouter}
}
