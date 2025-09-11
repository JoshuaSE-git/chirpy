package main

import (
	"net/http"

	"github.com/JoshuaSE-git/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing Authorization header"))
		return
	}
	err = cfg.db.RevokeToken(r.Context(), refreshTokenString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid token"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
