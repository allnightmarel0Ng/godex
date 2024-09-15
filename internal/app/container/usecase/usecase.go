package usecase

import (
	"errors"
	"strings"

	"github.com/allnightmarel0Ng/godex/internal/app/container/repository"
	"github.com/allnightmarel0Ng/godex/internal/domain/model"
)

type ContainerUseCase interface {
	ProcessNewFunction(function string) error
	ProcessGetFunction(signature string) ([]model.FunctionMetadata, error)
}

type containerUseCase struct {
	repo repository.ContainerRepository
}

func NewContainerUseCase(repo repository.ContainerRepository) ContainerUseCase {
	return &containerUseCase{
		repo: repo,
	}
}

func (c *containerUseCase) ProcessNewFunction(function string) error {
	tokens := strings.Split(function, " ")
	if len(tokens) != 6 {
		return errors.New("invalid function")
	}

	metadata := model.FunctionMetadata{
		Name:      tokens[0],
		Signature: tokens[1],
		Comment:   tokens[2],
		File: model.FileMetadata{
			Name: tokens[3],
			Package: model.PackageMetadata{
				Name: tokens[4],
				Link: tokens[5],
			},
		},
	}

	return c.repo.InsertFunction(metadata)
}

func (c *containerUseCase) ProcessGetFunction(signature string) ([]model.FunctionMetadata, error) {
	strings.TrimSpace(signature)
	return c.repo.GetFunctionBySignature(signature)
}
