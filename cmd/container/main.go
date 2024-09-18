package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/allnightmarel0Ng/godex/internal/app/container/handler"
	pb "github.com/allnightmarel0Ng/godex/internal/app/container/proto"
	"github.com/allnightmarel0Ng/godex/internal/app/container/repository"
	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load config")
	}

	db, err := postgres.NewDatabase(context.Background(), fmt.Sprintf("postgresql://%s:%s@postgres:%s/%s?sslmode=disable", conf.PostgresUser, conf.PostgresPassword, conf.PostgresPort, conf.PostgresName))
	if err != nil {
		log.Fatalf("unable to connect to the database: %s", err.Error())
	}
	defer db.Close()

	repo := repository.NewContainerRepository(db)
	useCase := usecase.NewContainerUseCase(repo)

	consumer, err := kafka.NewConsumer("kafka:9092", "functions")
	if err != nil {
		log.Fatalf("unable to create kafka consumer: %s", err.Error())
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{"functions"})
	if err != nil {
		log.Fatalf("unable to subscribe kafka consumer on topic: %s", err.Error())
	}

	kafkaHandler := handler.NewContainerKafkaHandler(consumer, useCase)
	go kafkaHandler.Handle()

	listener, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("unable to listen on port 5001: %s", err.Error())
	}

	server := grpc.NewServer()
	pb.RegisterContainerServer(server, &handler.ContainerGRPCHandler{UseCase: useCase})

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("unable to serve on port 5001: %s", err.Error())
	}

}
