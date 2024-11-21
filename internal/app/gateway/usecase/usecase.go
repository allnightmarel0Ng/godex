package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/allnightmarel0Ng/godex/internal/app/gateway/repository"
)

type GatewayUseCase interface {
	Store(requestBody []byte) ([]byte, error)
	Find(signature string) ([]byte, error)
}

type gatewayUseCase struct {
	repo       repository.GatewayRepository
	parserPort string
}

func NewGatewayUseCase(repo repository.GatewayRepository, parserPort string) GatewayUseCase {
	return &gatewayUseCase{
		repo:       repo,
		parserPort: parserPort,
	}
}

func (e *gatewayUseCase) Store(requestBody []byte) ([]byte, error) {
	response, err := http.Post(fmt.Sprintf("http://parser:%s/", e.parserPort), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	return body, err
}

func (e *gatewayUseCase) Find(signature string) ([]byte, error) {
	signature = strings.ReplaceAll(signature, " ", "")

	functions, err := e.repo.GetFunctionBySignature(signature)
	if err != nil {
		return nil, err
	}

	return json.Marshal(functions)
}
