package handler

import (
	"context"
	"net/http"
)

const (
	UserId = "UserId"
)

func (h *Handler) authorizationOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		userId, err := h.services.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if _, err := h.services.GetUserById(userId); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserId, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
