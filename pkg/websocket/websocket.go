package websocket

import "net/http"

type JsonWriter interface {
	WriteJson(userId int, message interface{})
	UpgradeConnection(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	JsonWriter
}

func NewServer() *Server {
	return &Server{
		JsonWriter: NewGorillaServer(),
	}
}
