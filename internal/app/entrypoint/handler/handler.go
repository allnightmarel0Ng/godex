package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/usecase"
	"github.com/allnightmarel0Ng/godex/internal/logger"
)

type EntrypointHandler struct {
	useCase usecase.EntrypointUseCase
}

func NewEntrypointHandler(useCase usecase.EntrypointUseCase) EntrypointHandler {
	return EntrypointHandler{
		useCase: useCase,
	}
}

func (e *EntrypointHandler) HandleStore(w http.ResponseWriter, r *http.Request) {
	logger.Debug("HandleStore: start")
	defer logger.Debug("HandleStore: end")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Warning("error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var link struct {
		Link string `json:"link"`
	}
	err = json.Unmarshal(body, &link)
	if err != nil {
		logger.Warning("error parsing JSON")
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	code, message, err := e.useCase.Store(link.Link)
	if err != nil {
		logger.Warning("error downloading file from link")
		http.Error(w, "Error downloading file from link", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(int(code))
	w.Write([]byte(message))
}

func (e *EntrypointHandler) HandleFind(w http.ResponseWriter, r *http.Request) {
	logger.Debug("HandleFind: start")
	defer logger.Debug("HandleFind: end")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Warning("error reading request body")
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var signature struct {
		Signature string `json:"signature"`
	}
	err = json.Unmarshal(body, &signature)
	if err != nil {
		logger.Warning("error parsing JSON")
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	functions, err := e.useCase.Find(signature.Signature)
	if err != nil || functions == nil {
		logger.Warning("error finding signature")
		http.Error(w, "Error finding signature", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(functions)
}
