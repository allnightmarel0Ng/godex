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

func NewGatewayHandler(useCase usecase.GatewayUseCase) GatewayHandler {
	return GatewayHandler{
		useCase: useCase,
	}
}

func (e *GatewayHandler) HandleStore(w http.ResponseWriter, r *http.Request) {
	logger.Debug("HandleStore: start")
	defer logger.Debug("HandleStore: end")

	if r.Header.Get("Content-Type") != "application/json" {
		logger.Warning("wrong content type")
		http.Error(w, "wrong content type: should be application/json", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Warning("error reading request body")
		http.Error(w, "error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var link struct {
		Link string `json:"link"`
	}
	err = json.Unmarshal(body, &link)
	if err != nil {
		logger.Warning("error parsing JSON")
		http.Error(w, "error parsing JSON", http.StatusBadRequest)
		return
	}

	code, message, err := e.useCase.Store(link.Link)
	if err != nil {
		logger.Warning(message)
		http.Error(w, "unexpected error", http.StatusInternalServerError)
		return
	}

	if code != http.StatusOK {
		logger.Warning(message)
		http.Error(w, message, int(code))
		return
	}

	w.WriteHeader(int(code))
	w.Write([]byte(message))
}

func (e *GatewayHandler) HandleFind(w http.ResponseWriter, r *http.Request) {
	logger.Debug("HandleFind: start")
	defer logger.Debug("HandleFind: end")

	if r.Header.Get("Content-Type") != "application/json" {
		logger.Warning("wrong content type")
		http.Error(w, "wrong content type: should be application/json", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Warning("error reading request body")
		http.Error(w, "error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var signature struct {
		Signature string `json:"signature"`
	}
	err = json.Unmarshal(body, &signature)
	if err != nil {
		logger.Warning("error parsing JSON")
		http.Error(w, "error parsing JSON", http.StatusBadRequest)
		return
	}

	functions, err := e.useCase.Find(signature.Signature)
	if err != nil || functions == nil {
		logger.Warning("error finding signature")
		http.Error(w, "error finding signature", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(functions)
}
