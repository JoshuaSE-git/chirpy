package main

import (
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("reset only allowed in dev enviornment"))
		return
	}

	err := cfg.db.ResetUsers(r.Context())
	if err != nil {
		log.Printf("failed to reset database: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to reset database: " + err.Error()))
		return
	}

	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hits reset to 0; users table reset"))
}
