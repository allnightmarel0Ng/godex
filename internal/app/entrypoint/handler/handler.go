package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/entrypoint/usecase"
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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var link struct {
		Link string `json:"link"`
	}
	err = json.Unmarshal(body, &link)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
	}

	code, message, err := e.useCase.Store(link.Link)
	if err != nil {
		http.Error(w, "Error downloading file from link", http.StatusInternalServerError)
	}

	w.WriteHeader(int(code))
	w.Write([]byte(message))
}

func (e *EntrypointHandler) HandleFind(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var signature struct {
		Signature string `json:"signature"`
	}
	err = json.Unmarshal(body, &signature)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
	}

	functions, err := e.useCase.Find(signature.Signature)
	if err != nil {
		http.Error(w, "Error finding signature", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(functions)
}
