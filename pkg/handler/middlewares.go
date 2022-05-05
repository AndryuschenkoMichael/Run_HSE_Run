package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	UserId = "UserId"
)

func (h *Handler) authorizationOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		tokenString := strings.Split(bearerToken, " ")
		if len(tokenString) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if tokenString[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId, err := h.services.ParseToken(tokenString[1])
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
