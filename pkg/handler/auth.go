package handler

import (
	"Run_Hse_Run/pkg/logger"
	"Run_Hse_Run/pkg/model"
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
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if email.Email == "" {
		logger.WarningLogger.Println("invalid email")
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := h.services.SendEmail(email.Email); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var auth struct {
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
		Image    int    `json:"image"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&auth); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if auth.Nickname == "" || auth.Email == "" {
		logger.WarningLogger.Println("invalid email or nickname")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.services.GetUser(auth.Email)
	if err == nil {
		logger.WarningLogger.Println("user already exist")
		w.WriteHeader(http.StatusConflict)
		return
	}

	user := model.User{
		Email:    auth.Email,
		Nickname: auth.Nickname,
		Image:    auth.Image,
	}

	id, err := h.services.CreateUser(user)

	if err != nil {
		logger.WarningLogger.Println(err)
		if err.Error() == service.NicknameError {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	user.Id = id

	token, err := h.services.GenerateToken(auth.Email)
	if err != nil {
		logger.WarningLogger.Println(err)
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
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if auth.Code == 0 || auth.Email == "" {
		logger.WarningLogger.Println("invalid email or code")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	service.Mu.Lock()
	code, ok := service.Codes[auth.Email]
	service.Mu.Unlock()

	if !ok {
		logger.WarningLogger.Println("email didn't added")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if code != auth.Code {
		logger.WarningLogger.Println("incorrect code")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := h.services.GetUser(auth.Email)
	if err != nil {
		logger.WarningLogger.Println("user doesn't exist in db")
		w.WriteHeader(http.StatusOK)
		return
	}

	token, err := h.services.GenerateToken(auth.Email)
	if err != nil {
		logger.WarningLogger.Println(err)
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
