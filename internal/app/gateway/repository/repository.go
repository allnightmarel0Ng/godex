package repository

import (
	"github.com/allnightmarel0Ng/godex/internal/domain/model"
	"github.com/allnightmarel0Ng/godex/internal/domain/repository"
)

type GatewayRepository interface {
	GetFunctionBySignature(signature string) ([]model.FunctionMetadata, error)
}

type gatewayRepository struct {
	functions repository.FunctionRepository
}

func NewGatewayRepositiry(functionRepository repository.FunctionRepository) GatewayRepository {
	return &gatewayRepository{
		functions: functionRepository,
	}
}

func (c *gatewayRepository) GetFunctionBySignature(signature string) ([]model.FunctionMetadata, error) {
	return c.functions.GetFunctionsBySignature(signature)
}
