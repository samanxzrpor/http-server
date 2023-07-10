package main

import (
	"net/http"

	"github.com/samanxzrpor/http-server/pkg/config"
	"github.com/samanxzrpor/http-server/pkg/server"
)

func main() {

	config := config.LoadConfigOrPanic()
	server := server.NewHttpServer(config.Server)
	server.HandlerFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	server.Start()
}
