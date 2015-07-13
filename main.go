package main

import (
	"fmt"
	"main/httpserver"
	"net"
)

type HelloHandler struct {
	name string
}

func (h HelloHandler) String() string {
	return "HelloHandler"
}

func generateHeader(contentLength int) string {
	head := "HTTP/1.1 200 OK\r\n"
	head += fmt.Sprintf("Content-length: %d\r\n", contentLength)
	return head
}

func (h HelloHandler) Handle(r *httpserver.HttpRequest, c net.Conn) error {
	content := "Hello, " + h.name
	header := generateHeader(len(content))
	message := header + "\r\n" + content + "\r\n"
	fmt.Fprint(c, message)
	return nil
}

func main() {
	server := httpserver.NewServer()
	server.AddHandle("/", HelloHandler{"Bob"})
	fmt.Printf("%v\n", server.Handlers["/"].String())
	server.Run(":8080")
}
