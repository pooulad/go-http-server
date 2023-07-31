package main

import (
	"net/http"
	"github.com/pooulad/go-http-server/pkg/config"
	"github.com/pooulad/go-http-server/pkg/server"
)

func main(){
	cnf := config.LoadConfigOrPanic()
	server := server.NewHttpServer(cnf.Server)
	server.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	server.Start()
}