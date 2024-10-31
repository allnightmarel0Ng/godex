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

type errorResponse struct {
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
	json.NewEncoder(w).Encode(errorResponse{
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

	code, message, err := e.useCase.Store(link.Link)
	if err != nil {
		e.sendError(w, http.StatusBadRequest, "unexpected error")
		logger.Warning(message)
		return
	}

	if code != http.StatusOK {
		e.sendError(w, int(code), message)
		logger.Warning(message)
		return
	}

	w.WriteHeader(int(code))
	w.Write([]byte(message))
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

	functions, err := e.useCase.Find(signature.Signature)
	if err != nil || functions == nil {
		e.sendError(w, http.StatusNotFound, "error finding signature")
		logger.Warning("error finding signature")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(functions)
}
