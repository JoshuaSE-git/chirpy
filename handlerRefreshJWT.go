package main

import (
	"net/http"
	"time"

	"github.com/JoshuaSE-git/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefreshJWT(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}
	refreshTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "missing authorization header", err)
		return
	}
	refreshToken, err := cfg.db.GetUserFromRefreshToken(r.Context(), refreshTokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid token", err)
		return
	}
	if time.Now().After(refreshToken.ExpiresAt) {
		respondWithError(w, http.StatusUnauthorized, "expired token", err)
		return
	}
	if refreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "revoked token", err)
		return
	}
	jwtTokenString, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret, DefaultJWTExpiration)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create jwt", err)
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		Token: jwtTokenString,
	})
}
