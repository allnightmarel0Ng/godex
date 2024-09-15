package handler

import (
	"context"

	pb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
)

type ContainerGRPCHandler struct {
	UseCase usecase.ContainerUseCase
	pb.UnimplementedContainerServer
}

func (c *ContainerGRPCHandler) Find(ctx context.Context, in *pb.SignatureRequest) (*pb.FunctionsResponse, error) {
	signature := in.GetSignature()
	metadatas, err := c.UseCase.ProcessGetFunction(signature)
	if err != nil {
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
