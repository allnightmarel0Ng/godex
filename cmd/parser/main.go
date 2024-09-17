package main

import (
	"log"
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
		log.Fatalf("unable to load config")
	}

	listener, err := net.Listen("tcp", ":"+conf.ParserPort)
	if err != nil {
		log.Fatalf("unable to create the listener")
	}
	defer listener.Close()

	producer, err := kafka.NewProducer("kafka:" + conf.KafkaBroker)
	if err != nil {
		log.Fatalf("unable to create a kafka producer")
	}
	defer producer.Close()

	useCase := usecase.NewParserUseCase(producer, conf.WhiteList)

	s := grpc.NewServer()
	pb.RegisterParserServer(s, &handler.ParserHandler{UseCase: useCase})
	defer s.GracefulStop()

	if err = s.Serve(listener); err != nil {
		log.Fatalf("unable to start gRPC server: %s", err.Error())
	}
}
