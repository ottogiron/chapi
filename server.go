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

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Chapi")
}

func main() {

	connectionString := ":" + os.Getenv("PORT")
	s := server.NewServer()
	s.Register(&HelloController{})
	s.Run(connectionString)

}
