package handler

import (
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) sendEmail(w http.ResponseWriter, r *http.Request) {
	var email struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.services.SendEmail(email.Email); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var auth struct {
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&auth); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.services.GetUser(auth.Email)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	user := model.User{
		Email:    auth.Email,
		Nickname: auth.Nickname,
	}

	id, err := h.services.CreateUser(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	user.Id = id

	token, err := h.services.GenerateToken(auth.Email)
	if err != nil {
		fmt.Println("kek")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userToken := struct {
		Token string     `json:"token"`
		User  model.User `json:"user"`
	}{
		Token: token,
		User:  user,
	}

	tokenString, _ := json.Marshal(userToken)

	w.WriteHeader(http.StatusCreated)
	w.Write(tokenString)
}

func (h *Handler) checkAuth(w http.ResponseWriter, r *http.Request) {
	var auth struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&auth); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	service.Mu.Lock()
	code, ok := service.Codes[auth.Email]
	service.Mu.Unlock()

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if code != auth.Code {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := h.services.GetUser(auth.Email)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	token, err := h.services.GenerateToken(auth.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userToken := struct {
		Token string     `json:"token"`
		User  model.User `json:"user"`
	}{
		Token: token,
		User:  user,
	}

	tokenString, _ := json.Marshal(userToken)

	w.WriteHeader(http.StatusCreated)
	w.Write(tokenString)
}
