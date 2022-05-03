package handler

import (
	"context"
	"net/http"
)

func (h *Handler) authorizationOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		userId, err := h.services.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "UserId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
