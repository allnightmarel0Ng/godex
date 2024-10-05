package main

import (
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"net"

	"github.com/allnightmarel0Ng/godex/internal/app/parser/handler"
	pb "github.com/allnightmarel0Ng/godex/internal/app/parser/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	listener, err := net.Listen("tcp", ":"+conf.ParserPort)
	if err != nil {
		logger.Error("unable to create the listener: %s", err.Error())
	}
	defer listener.Close()

	producer, err := kafka.NewProducer("kafka:" + conf.KafkaBroker)
	if err != nil {
		logger.Error("unable to create a kafka producer: %s", err.Error())
	}
	defer producer.Close()

	useCase := usecase.NewParserUseCase(producer, conf.WhiteList)

	s := grpc.NewServer()
	pb.RegisterParserServer(s, &handler.ParserHandler{UseCase: useCase})
	defer s.GracefulStop()

	if err = s.Serve(listener); err != nil {
		logger.Error("unable to start gRPC server: %s", err.Error())
	}
}
