package usecase

import (
	"encoding/json"

	"github.com/allnightmarel0Ng/godex/internal/app/container/repository"
	"github.com/allnightmarel0Ng/godex/internal/domain/model"
)

type ContainerUseCase interface {
	ProcessNewFunction(function []byte) error
}

type containerUseCase struct {
	repo repository.ContainerRepository
}

func NewContainerUseCase(repo repository.ContainerRepository) ContainerUseCase {
	return &containerUseCase{
		repo: repo,
	}
}

func (c *containerUseCase) ProcessNewFunction(function []byte) error {
	var metadata model.FunctionMetadata
	err := json.Unmarshal(function, &metadata)
	if err != nil {
		return err
	}
	return c.repo.InsertFunction(metadata)
}
