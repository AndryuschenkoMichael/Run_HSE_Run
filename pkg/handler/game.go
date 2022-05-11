package handler

import (
	"Run_Hse_Run/pkg/logger"
	"encoding/json"
	"net/http"
)

func (h *Handler) getRoomByCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	rooms, err := h.services.GetRoomByCodePattern(code)
	if err != nil {
		logger.WarningLogger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	roomsJson, _ := json.Marshal(rooms)
	w.WriteHeader(http.StatusOK)
	w.Write(roomsJson)
}
