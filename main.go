package main

import (
	"net/http"
)

const addr = ":8080"

func main() {
	srvmux := http.NewServeMux()
	server := http.Server{
		Addr:    addr,
		Handler: srvmux,
	}
	server.ListenAndServe()
}
