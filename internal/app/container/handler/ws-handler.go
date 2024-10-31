package handler

import (
	"encoding/json"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/gorilla/websocket"
)

type ContainerWebSocketHandler struct {
	upgrader websocket.Upgrader
	useCase usecase.ContainerUseCase
}

func NewContainerWebSocketHandler(useCase usecase.ContainerUseCase) ContainerWebSocketHandler {
	return ContainerWebSocketHandler{
		upgrader: websocket.Upgrader{},
		useCase: useCase,
	}
}

func send(conn *websocket.Conn, message []byte) {
	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		logger.Warning("write message error")
	}
}

func (c ContainerWebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("ServeHTTP: start")
	defer logger.Debug("ServeHTTP: end")

	conn, err := c.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Warning("unable to set ws connection: %s", err.Error())
		return
	}
	defer conn.Close()

	_, signature, err := conn.ReadMessage()
	if err != nil {
		logger.Info("disconnected")
		return
	}

	functions, err := c.useCase.ProcessGetFunction(string(signature))
	if err != nil {
		logger.Warning("unable to retrive functions from database: %s", err.Error())
		send(conn, []byte("DB_ERROR"))
		return
	}

	if functions == nil {
		logger.Info("didn't found anything")
		send(conn, []byte("NOT_FOUND"))
		return
	}

	bytes, _ := json.Marshal(functions)
	send(conn, bytes)
}