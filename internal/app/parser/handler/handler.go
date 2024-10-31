package handler

import (
	"context"
	"fmt"
	"io"
	"net/http"

	pb "github.com/allnightmarel0Ng/godex/internal/app/parser/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
	"github.com/allnightmarel0Ng/godex/internal/logger"
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
		return nil, fmt.Errorf("%s", response.Status)
	}

	return io.ReadAll(response.Body)
}

func (p *ParserHandler) Download(ctx context.Context, in *pb.LinkRequest) (*pb.StatusReply, error) {
	logger.Debug("Download: start")
	defer logger.Debug("Download: end")

	url := in.GetLink()
	fileName, packageName, link, err := p.UseCase.ParseUrl(url)
	if err != nil {
		message := fmt.Sprintf("invalid link: %s", err.Error())
		logger.Warning(message)
		return &pb.StatusReply{
			Status:  http.StatusBadRequest,
			Message: message,
		}, nil
	}

	bytes, err := p.fetchFile(url)
	if err != nil {
		message := fmt.Sprintf("unable to fetch the data from link: %s", err.Error())
		logger.Warning(message)
		return &pb.StatusReply{
			Status:  http.StatusNotFound,
			Message: message,
		}, nil
	}

	functions, err := p.UseCase.ExtractFunctions(bytes, fileName, packageName, link)
	if err != nil {
		message := fmt.Sprintf("unable to get functions from file: %s", err.Error())
		logger.Warning(message)
		return &pb.StatusReply{
			Status:  http.StatusInternalServerError,
			Message: message,
		}, nil
	}

	for _, function := range functions {
		err = p.UseCase.ProduceMessage(function)
		if err != nil {
			logger.Warning("producer error: %s", err.Error())
		}
	}

	return &pb.StatusReply{
		Status:  http.StatusOK,
		Message: "success",
	}, nil
}
