package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.services.GetUserById(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJson, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}

func (h *Handler) getUserByNickname(w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query().Get("nickname")

	users, err := h.services.GetUsersByNicknamePattern(nickname)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usersJson, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	w.Write(usersJson)
}
