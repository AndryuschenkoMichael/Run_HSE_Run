package handler

import (
	"Run_Hse_Run/pkg/logger"
	"encoding/json"
	"net/http"
)

func (h *Handler) sendTime(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var gameInfo struct {
		GameId int `json:"game_id"`
		Time   int `json:"time"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&gameInfo); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.services.UpdateTime(gameInfo.GameId, userId, gameInfo.Time); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	go h.services.SendResult(gameInfo.GameId, userId, gameInfo.Time)

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getRoomByCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	rooms, err := h.services.GetRoomByCodePattern(code, 1)
	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	roomsJson, _ := json.Marshal(rooms)
	w.WriteHeader(http.StatusOK)
	w.Write(roomsJson)
}

func (h *Handler) putInQueue(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var roomId struct {
		RoomId int `json:"room_id"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&roomId); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	h.services.AddUser(userId, roomId.RoomId)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteFromQueue(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.services.Cancel(userId)

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) addCall(w http.ResponseWriter, r *http.Request) {
	userIdFirst, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var call struct {
		RoomIdFirst  int `json:"room_id_first"`
		UserIdSecond int `json:"user_id_second"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&call); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	game, err := h.services.AddCall(userIdFirst, call.UserIdSecond, call.RoomIdFirst)
	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if game.UserIdFirst == -1 {
		w.WriteHeader(http.StatusCreated)
		return
	}

	err = h.services.SendGame(game)

	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteCall(w http.ResponseWriter, r *http.Request) {
	userIdFirst, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var call struct {
		UserIdSecond int `json:"user_id_second"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&call); err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.services.DeleteCall(userIdFirst, call.UserIdSecond)
	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
