package usecase

import (
	"context"

	containerpb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	parserpb "github.com/allnightmarel0Ng/godex/internal/app/parser/proto"
	"github.com/allnightmarel0Ng/godex/internal/domain/model"
)

type EntrypointUseCase interface {
	Store(link string) (int32, string, error)
	Find(signature string) ([]model.FunctionMetadata, error)
}

type entrypointUseCase struct {
	toParser    parserpb.ParserClient
	toContainer containerpb.ContainerClient
}

func NewEntrypointUseCase(ParserClient parserpb.ParserClient, ContainerClient containerpb.ContainerClient) EntrypointUseCase {
	return &entrypointUseCase{
		toParser:    ParserClient,
		toContainer: ContainerClient,
	}
}

func (e *entrypointUseCase) Store(link string) (int32, string, error) {
	response, err := e.toParser.Download(context.Background(), &parserpb.LinkRequest{Link: link})
	if err != nil {
		return 0, "", err
	}

	return response.GetStatus(), response.GetMessage(), nil
}

func (e *entrypointUseCase) Find(signature string) ([]model.FunctionMetadata, error) {
	response, err := e.toContainer.Find(context.Background(), &containerpb.SignatureRequest{Signature: signature})
	if err != nil {
		return nil, err
	}

	var result []model.FunctionMetadata
	for _, function := range response.Functions {
		result = append(result, model.FunctionMetadata{
			Name:      function.FunctionName,
			Signature: function.FunctionSignature,
			Comment:   function.FunctionComment,
			File: model.FileMetadata{
				Name: function.FileName,
				Package: model.PackageMetadata{
					Name: function.PackageName,
					Link: function.PackageLink,
				},
			},
		})
	}

	return result, nil
}
