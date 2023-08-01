package main

import (
	"context"
	"net/http"

	"github.com/pooulad/go-http-server/internal/db"
	"github.com/pooulad/go-http-server/internal/db/postgres"
	"github.com/pooulad/go-http-server/pkg/config"
	"github.com/pooulad/go-http-server/pkg/server"
)

func main() {
	cnf := config.LoadConfigOrPanic()
	pg,err := postgres.NewPostgres(cnf.Postgres)
	if err != nil {
		panic(err)
	}
	if err := db.MigrateOrPanic(context.Background(),pg);err != nil {
		panic(err)
	}
	server := server.NewHttpServer(cnf.Server)
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	server.Start()
}
