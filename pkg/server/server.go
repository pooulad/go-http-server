package server

import (
	"fmt"
	"net/http"
	"github.com/pooulad/go-http-server/pkg/config"
)

type HttpServer struct {
	server http.Server
	mux    *http.ServeMux
}

func NewHttpServer(cnf config.Config) *HttpServer {
	fmt.Println("server started at port:",cnf.Port)
	return &HttpServer{
		server: http.Server{
			Addr: fmt.Sprintf("%s:%d", cnf.Host, cnf.Port),
		},
		mux: http.NewServeMux(),
	}
}

func (s *HttpServer) HandleFunc(route string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(route, handler)
}

func (s *HttpServer) Start() error {
	s.server.Handler = s.mux
	return s.server.ListenAndServe()
}
