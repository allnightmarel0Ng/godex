package main

import (
	"log"
	"net"

	"github.com/allnightmarel0Ng/godex/internal/app/container/handler"
	pb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/container/repository"
	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewDatabase("user=admin dbname=godex password=admin sslmode=disable")
	if err != nil {
		log.Fatalf("unable to connect to the database: %s", err.Error())
	}

	repo := repository.NewContainerRepository(db)
	useCase := usecase.NewContainerUseCase(repo)

	consumer, err := kafka.NewConsumer("localhost:9092", "foo")
	if err != nil {
		log.Fatalf("unable to create kafka consumer: %s", err.Error())
	}

	err = consumer.SubscribeTopics([]string{"functions"})
	if err != nil {
		log.Fatalf("unable to subscribe kafka consumer on topic: %s", err.Error())
	}

	kafkaHandler := handler.NewContainerKafkaHandler(consumer, useCase)
	go kafkaHandler.Handle()

	listener, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("unable to listen on port 5051: %s", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterContainerServer(server, &handler.ContainerGRPCHandler{UseCase: useCase})

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("unable to serve on port 5051: %s", err.Error())
	}

}
