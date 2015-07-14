package main

import (
	"fmt"
	"main/httpserver"
)

type HelloHandler struct {
	name string
}

func (h HelloHandler) String() string {
	return "HelloHandler"
}

func (h HelloHandler) Handle(req *httpserver.HttpRequest, res *httpserver.HttpResponse) {
	body := "Hello, " + h.name
	res.SetStatus("OK")
	fmt.Fprint(res, body)
	fmt.Fprint(res, "\nThis is another line")
}

func main() {
	server := httpserver.NewServer()
	server.AddHandle("/", HelloHandler{"Bob"})
	server.AddHandle("/frank", HelloHandler{"Frank"})
	server.AddHandle("/russel", HelloHandler{"Russel"})
	fmt.Printf("%v\n", server.Handlers["/"].String())
	server.Run(":8080")
}
