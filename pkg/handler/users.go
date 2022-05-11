package handler

import (
	"Run_Hse_Run/pkg/logger"
	"Run_Hse_Run/pkg/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) getUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.services.GetUserById(userId)
	if err != nil {
		logger.WarningLogger.Println(err)
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
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	usersJson, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	w.Write(usersJson)
}

func (h *Handler) changeNickname(w http.ResponseWriter, r *http.Request) {
	var newNickname struct {
		Nickname string `json:"nickname"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newNickname); err != nil {
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

	err := h.services.RenameUser(userId, newNickname.Nickname)

	if err != nil {
		logger.WarningLogger.Println(err)
		if err.Error() == service.NicknameError {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) changeProfileImage(w http.ResponseWriter, r *http.Request) {
	var profileImage struct {
		Image int `json:"image"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&profileImage); err != nil {
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

	err := h.services.ChangeProfileImage(userId, profileImage.Image)

	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.services.GetUserById(userId)
	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJson, _ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
}
