package websocket

import (
	"Run_Hse_Run/pkg/logger"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const timeOut = 3 * time.Minute

type GorillaServer struct {
	clients map[int]*websocket.Conn
}

func (g *GorillaServer) WriteJson(userId int, message interface{}) {
	timer := time.NewTimer(timeOut)
	for {
		select {
		case <-timer.C:
			logger.WarningLogger.Println("time out of write json")
			return
		default:
			connection, ok := g.clients[userId]
			if !ok {
				logger.WarningLogger.Println("connection doesn't exist")
			} else {
				if err := connection.WriteJSON(message); err == nil {
					return
				} else {
					logger.WarningLogger.Printf("connection was lost: %s", err.Error())
				}
			}
		}
	}
}

func (g *GorillaServer) UpgradeConnection(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("UserId").(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		return
	}

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.WarningLogger.Printf("can't upgrade connection: %s", err.Error())
	}

	defer connection.Close()

	if con, ok := g.clients[userId]; ok {
		err := con.Close()
		if err != nil {
			logger.WarningLogger.Printf("can't close websocket: %s", err.Error())
		}
	}

	g.clients[userId] = connection
	defer delete(g.clients, userId)

	for {
		mt, _, err := connection.ReadMessage()

		logger.WarningLogger.Printf("receive message type: %d", mt)

		if err != nil || mt == websocket.CloseMessage {
			break
		}

		if mt == websocket.PingMessage {
			err := connection.WriteMessage(websocket.PongMessage, []byte{})
			if err != nil {
				logger.WarningLogger.Printf("connection lost: %s", err.Error())
				break
			}
		}
	}
}

func NewGorillaServer() *GorillaServer {
	return &GorillaServer{
		clients: make(map[int]*websocket.Conn),
	}
}
