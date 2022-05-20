package websocket

import (
	"Run_Hse_Run/pkg/logger"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	timeOut        = 3 * time.Minute
	timeQueryAgain = time.Second
	UserId         = "UserId"
)

type GorillaServer struct {
	sync.RWMutex
	clients map[int]*websocket.Conn
}

func (g *GorillaServer) WriteJson(userId int, message interface{}) {
	timer := time.NewTimer(timeOut)
	ticker := time.NewTicker(timeQueryAgain)
	for {
		select {
		case <-timer.C:
			logger.WarningLogger.Println("time out of write json")
			return
		case <-ticker.C:
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
	userId, ok := r.Context().Value(UserId).(int)
	if !ok {
		logger.WarningLogger.Println("invalid context")
		return
	}

	logger.WarningLogger.Printf("user with id = %d try to connect by web socket", userId)

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.WarningLogger.Printf("can't upgrade connection: %s", err.Error())
	}

	g.Lock()
	if con, ok := g.clients[userId]; ok {
		err := con.Close()
		if err != nil {
			logger.WarningLogger.Printf("can't close websocket: %s", err.Error())
		}
	}

	g.clients[userId] = connection
	g.Unlock()

	for {
		g.RLock()
		con, ok := g.clients[userId]
		if !ok {
			g.RUnlock()
			break
		}
		mt, _, err := con.ReadMessage()
		g.RUnlock()

		logger.WarningLogger.Printf("receive message type: %d", mt)

		if err != nil || mt == websocket.CloseMessage {
			break
		}

		if mt == websocket.PingMessage {
			g.RLock()
			con, ok := g.clients[userId]
			if !ok {
				g.RUnlock()
				break
			}

			err := con.WriteMessage(websocket.PongMessage, []byte{})
			if err != nil {
				logger.WarningLogger.Printf("connection lost: %s", err.Error())
				g.RUnlock()
				break
			}

			g.RUnlock()
		}
	}
}

func NewGorillaServer() *GorillaServer {
	return &GorillaServer{
		clients: make(map[int]*websocket.Conn),
	}
}
