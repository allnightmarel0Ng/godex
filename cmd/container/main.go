package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/container/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/container/repository"
	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	domainRepository "github.com/allnightmarel0Ng/godex/internal/domain/repository"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/postgres"
	"github.com/allnightmarel0Ng/godex/internal/logger"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	db, err := postgres.NewDatabase(context.Background(), fmt.Sprintf("postgresql://%s:%s@postgres:%s/%s?sslmode=disable", conf.PostgresUser, conf.PostgresPassword, conf.PostgresPort, conf.PostgresName))
	if err != nil {
		logger.Error("unable to connect to the database: %s", err.Error())
	}
	defer db.Close()

	repo := repository.NewContainerRepository(
		domainRepository.NewFunctionRepository(db),
		domainRepository.NewFileRepository(db),
		domainRepository.NewPackageRepository(db))
	useCase := usecase.NewContainerUseCase(repo)

	consumer, err := kafka.NewConsumer("kafka:9092", "functions")
	if err != nil {
		logger.Error("unable to create kafka consumer: %s", err.Error())
	}
	defer consumer.Close()

	err = consumer.SubscribeTopics([]string{"functions"})
	if err != nil {
		logger.Error("unable to subscribe kafka consumer on topic: %s", err.Error())
	}

	kafkaHandler := handler.NewContainerKafkaHandler(consumer, useCase)
	go kafkaHandler.Handle()

	wsHandler := handler.NewContainerWebSocketHandler(useCase)
	http.Handle("/", wsHandler)
	logger.Error("%s", http.ListenAndServe(":"+conf.ContainerPort, nil))
}
