package handler

import (
	"context"

	pb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/logger"
)

type ContainerGRPCHandler struct {
	UseCase usecase.ContainerUseCase
	pb.UnimplementedContainerServer
}

func (c *ContainerGRPCHandler) Find(ctx context.Context, in *pb.SignatureRequest) (*pb.FunctionsResponse, error) {
	logger.Debug("Find: start")
	defer logger.Debug("Find: end")

	signature := in.GetSignature()
	metadatas, err := c.UseCase.ProcessGetFunction(signature)
	if err != nil {
		logger.Warning("error while finding signature: %s", err.Error())
		return nil, err
	}

	var result []*pb.Function
	for _, metadata := range metadatas {
		result = append(result, &pb.Function{
			FunctionName:      metadata.Name,
			FunctionSignature: metadata.Signature,
			FunctionComment:   metadata.Comment,
			FileName:          metadata.File.Name,
			PackageName:       metadata.File.Package.Name,
			PackageLink:       metadata.File.Package.Link,
		})
	}

	return &pb.FunctionsResponse{Functions: result}, nil
}
