package main

import (
	"os"

	"github.com/ottogiron/chapi/server"
)

func main() {

	connectionString := ":" + os.Getenv("PORT")
	s := server.NewServer()
	//	s.Register(&controllers.HelloController{})
	s.Run(connectionString)

}
