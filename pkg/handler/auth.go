package handler

import (
	"Run_Hse_Run/pkg/service"
	"encoding/json"
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

	token, err := h.services.GenerateToken(auth.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userToken := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	tokenString, _ := json.Marshal(userToken)

	w.WriteHeader(http.StatusCreated)
	w.Write(tokenString)
}
