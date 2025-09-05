package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	const contentType = "text/plain; charset=utf-8"

	header := w.Header()
	header.Add("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
