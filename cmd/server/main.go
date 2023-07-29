package main

import (
	"net/http"
	"github.com/pooulad/go-http-server/pkg/server"
)

func main(){
	server := server.NewHttpServer("localhost",9090)
	server.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	server.Start()
}