package server

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
	server http.Server
	mux    http.ServeMux
}

func NewHttpServer(host string, port int) *HttpServer {
	return &HttpServer{
		server: http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
		mux: *http.NewServeMux(),
	}
}
