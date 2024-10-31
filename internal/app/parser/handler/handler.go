package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/gorilla/websocket"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ParserHandler struct {
	upgrader websocket.Upgrader
	useCase  usecase.ParserUseCase
}

func NewParserHandler(useCase usecase.ParserUseCase) ParserHandler {
	return ParserHandler{
		upgrader: websocket.Upgrader{},
		useCase:  useCase,
	}
}

func send(conn *websocket.Conn, code int, message string) {
	bytes, _ := json.Marshal(response{
		Code:    code,
		Message: message,
	})

	if err := conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
		logger.Warning("unable to send the message: %s", err.Error())
	}
}

func (p ParserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Debug("ServeHTTP: start")
	defer logger.Debug("ServeHTTP: end")

	conn, err := p.upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Warning("unable to set ws connection: %s", err.Error())
		return
	}
	defer conn.Close()

	_, url, err := conn.ReadMessage()
	if err != nil {
		logger.Info("disconnected")
		return
	}

	fileName, packageName, link, err := p.useCase.ParseUrl(string(url))
	if err != nil {
		message := fmt.Sprintf("invalid link: %s", err.Error())
		logger.Warning(message)
		send(conn, http.StatusBadRequest, message)
		return
	}

	bytes, err := p.useCase.FetchFile(string(url))
	if err != nil {
		message := fmt.Sprintf("unable to fetch the data from link: %s", err.Error())
		logger.Warning(message)
		send(conn, http.StatusNotFound, message)
		return
	}

	functions, err := p.useCase.ExtractFunctions(bytes, fileName, packageName, link)
	if err != nil {
		message := fmt.Sprintf("unable to get functions from file: %s", err.Error())
		logger.Warning(message)
		send(conn, http.StatusInternalServerError, message)
		return
	}

	for _, function := range functions {
		err = p.useCase.ProduceMessage(function)
		if err != nil {
			logger.Warning("producer error: %s", err.Error())
		}
	}

	send(conn, http.StatusOK, "success")
}
