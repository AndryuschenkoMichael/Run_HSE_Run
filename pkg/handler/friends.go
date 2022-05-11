package handler

import (
	"Run_Hse_Run/pkg/logger"
	"encoding/json"
	"net/http"
)

func (h *Handler) addFriend(w http.ResponseWriter, r *http.Request) {
	var userIdTo struct {
		UserId int `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userIdTo); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.services.AddFriend(userId, userIdTo.UserId)
	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) deleteFriend(w http.ResponseWriter, r *http.Request) {
	var userIdTo struct {
		UserId int `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userIdTo); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = h.services.DeleteFriend(userId, userIdTo.UserId)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) getFriends(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users, err := h.services.GetFriends(userId)

	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usersJson, _ := json.Marshal(users)

	w.WriteHeader(http.StatusOK)
	w.Write(usersJson)
}
