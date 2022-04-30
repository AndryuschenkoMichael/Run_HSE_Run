package handler

import (
	"context"
	"net/http"
)

func (h *Handler) authorizationOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		email := r.Header.Get("Email")

		token, err := h.services.GenerateToken(email)
		if err != nil || token != auth {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
