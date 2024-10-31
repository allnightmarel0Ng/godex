package usecase

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type GatewayUseCase interface {
	Store(link string) ([]byte, error)
	Find(signature string) ([]byte, error)
}

type gatewayUseCase struct {
	toParser    url.URL
	toContainer url.URL
}

func NewGatewayUseCase(parserURL url.URL, containerURL url.URL) GatewayUseCase {
	return &gatewayUseCase{
		toParser:    parserURL,
		toContainer: containerURL,
	}
}

func (e *gatewayUseCase) Store(link string) ([]byte, error) {
	parserConn, _, err := websocket.DefaultDialer.Dial(e.toParser.String(), nil)
	if err != nil {
		return nil, err
	}
	defer parserConn.Close()
	
	err = parserConn.WriteMessage(websocket.TextMessage, []byte(link))
	if err != nil {
		return nil, err
	}

	_, response, err := parserConn.ReadMessage()
	return response, err
}

func (e *gatewayUseCase) Find(signature string) ([]byte, error) {
	containerConn, _, err := websocket.DefaultDialer.Dial(e.toContainer.String(), nil)
	if err != nil {
		return nil, err
	}
	defer containerConn.Close()
	
	err = containerConn.WriteMessage(websocket.TextMessage, []byte(signature))
	if err != nil {
		return nil, err
	}

	_, response, err := containerConn.ReadMessage()
	return response, err
}
