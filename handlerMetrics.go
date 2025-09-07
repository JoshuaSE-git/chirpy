package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	const contentType = "text/html; charset=utf-8"
	const template = `<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>`

	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(template, cfg.fileserverHits.Load())))
}
