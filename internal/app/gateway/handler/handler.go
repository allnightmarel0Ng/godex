package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/gateway/usecase"
	"github.com/allnightmarel0Ng/godex/internal/logger"
)

type GatewayHandler struct {
	useCase usecase.GatewayUseCase
}

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewGatewayHandler(useCase usecase.GatewayUseCase) GatewayHandler {
	return GatewayHandler{
		useCase: useCase,
	}
}

func (e *GatewayHandler) sendError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response{
		Code:    code,
		Message: message,
	})
}

func (e *GatewayHandler) HandleStore(w http.ResponseWriter, r *http.Request) {
	logger.Debug("HandleStore: start")
	defer logger.Debug("HandleStore: end")

	if r.Header.Get("Content-Type") != "application/json" {
		e.sendError(w, http.StatusBadRequest, "wrong content type: should be application/json")
		logger.Warning("wrong content type")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		e.sendError(w, http.StatusInternalServerError, "error reading request body")
		logger.Warning("error reading request body")
		return
	}
	defer r.Body.Close()

	var link struct {
		Link string `json:"link"`
	}
	err = json.Unmarshal(body, &link)
	if err != nil {
		e.sendError(w, http.StatusBadRequest, "error parsing JSON")
		logger.Warning("error parsing JSON")
		return
	}

	payload, err := e.useCase.Store(link.Link)
	if err != nil {
		e.sendError(w, http.StatusInternalServerError, "unexpected error")
		logger.Warning("websocket error: %s", err.Error())
		return
	}

	var response response
	if json.Unmarshal(payload, &response) != nil {
		e.sendError(w, http.StatusInternalServerError, "unexpected error")
		logger.Warning("json unmarshalling error")
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	w.Write(payload)
}

func (e *GatewayHandler) HandleFind(w http.ResponseWriter, r *http.Request) {
	logger.Debug("HandleFind: start")
	defer logger.Debug("HandleFind: end")

	if r.Header.Get("Content-Type") != "application/json" {
		e.sendError(w, http.StatusBadRequest, "wrong content type: should be application/json")
		logger.Warning("wrong content type")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		e.sendError(w, http.StatusInternalServerError, "error reading request body")
		logger.Warning("error reading request body")
		return
	}
	defer r.Body.Close()

	var signature struct {
		Signature string `json:"signature"`
	}
	err = json.Unmarshal(body, &signature)
	if err != nil {
		e.sendError(w, http.StatusBadRequest, "error parsing JSON")
		logger.Warning("error parsing JSON")
		return
	}

	payload, err := e.useCase.Find(signature.Signature)
	if err != nil || payload == nil || string(payload) == "NOT_FOUND" {
		e.sendError(w, http.StatusNotFound, "error finding signature")
		logger.Warning("error finding signature")
		return
	}

	if string(payload) == "DB_ERROR" || string(payload) == "READ_ERROR" {
		e.sendError(w, http.StatusInternalServerError, "unexpected error")
		logger.Warning("container got an error: %s", string(payload))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
