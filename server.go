package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ottogiron/chapi/server"
)

type HelloController struct {
	*server.BasePlugin
}

func (helloController *HelloController) Register(server server.Server) {
	server.HandleFunc("/", handleHello).Methods("GET")
}

func (helloController *HelloController) Name() string {
	return "HelloController"
}

func (helloController *HelloController) Dependencies() []string {
	return []string{"HelloController"}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Chapi")
}

type HelloDep struct {
	*server.BasePlugin
}

func (helloDep *HelloDep) Name() string {
	return "HelloDep"
}

func (helloDep *HelloDep) Register(server server.Server) {
	server.HandleFunc("/hellodep", handleHello).Methods("GET")
}

func main() {

	connectionString := ":" + os.Getenv("PORT")
	s := server.NewServer()
	s.Register(&HelloController{})
	s.Register(&HelloDep{})
	s.Run(connectionString)

}
