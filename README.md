# Chapi Toy Web Framework
### Implemented for learning purposes


### Package server

```go
import "github.com/ottogiron/chapi/server"
```

#### Installation
```bash
go get github.com/ottogiron/chapi (didn't test this)
```

#### server.Server

Server is the application container.

**server.NewServer()**

Creates a new server

**Server.Register(plugin Plugin)**

Registers a new server plugin

**Server.Run(add String)**

Runs a server on the specified address

**Server.HandleFunc(path string, f func(http.ResponseWriter, *http.Request))  route**

Registers a new handler function.


***Example***

server.go
```go
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

	connectionString := ":8000"
	s := server.NewServer()
	s.Register(&HelloController{})
	s.Run(connectionString)

}
```
***Running the server***
```
go run server.go
```
