package handler

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	pb "github.com/allnightmarel0Ng/godex/internal/app/parser/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
)

type ParserHandler struct {
	UseCase usecase.ParserUseCase
	pb.UnimplementedParserServer
}

func (p *ParserHandler) fetchFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", response.Status)
	}

	return io.ReadAll(response.Body)
}

func (p *ParserHandler) Download(ctx context.Context, in *pb.LinkRequest) (*pb.StatusReply, error) {
	url := in.GetLink()
	bytes, err := p.fetchFile(url)
	if err != nil {
		return &pb.StatusReply{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
		}, fmt.Errorf("unable to fetch the data from link: %s", err.Error())
	}

	functions, err := p.UseCase.ExtractFunctions(bytes, url)
	if err != nil {
		return &pb.StatusReply{
			Status:  http.StatusUnprocessableEntity,
			Message: err.Error(),
		}, fmt.Errorf("unable to get functions from file: %s", err.Error())
	}

	for _, function := range functions {
		err = p.UseCase.ProduceMessage(function)
		if err != nil {
			log.Printf("producer error: %s", err.Error())
		}
	}

	return &pb.StatusReply{
		Status:  http.StatusOK,
		Message: "OK",
	}, nil
}
