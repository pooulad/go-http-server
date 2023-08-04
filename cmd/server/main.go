package main

import (
	"context"
	"github.com/pooulad/go-http-server/internal/db"
	"github.com/pooulad/go-http-server/internal/db/postgres"
	"github.com/pooulad/go-http-server/internal/handler"
	"github.com/pooulad/go-http-server/internal/repo"
	"github.com/pooulad/go-http-server/pkg/config"
	"github.com/pooulad/go-http-server/pkg/server"
)

func main() {
	cnf := config.LoadConfigOrPanic()
	pg, err := postgres.NewPostgres(cnf.Postgres)
	if err != nil {
		panic(err)
	}
	if err := db.MigrateOrPanic(context.Background(), pg); err != nil {
		panic(err)
	}

	dbRepo := repo.NewTrackRepo(pg)
	trackHandler := handler.NewTrackHandler(dbRepo)
	server := server.NewHttpServer(cnf.Server)
	server.HandleFunc("/track", trackHandler.Handle)
	server.Start()
}
