package main

import (
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/parser/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/godex/internal/logger"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		logger.Error("unable to load config: %s", err.Error())
	}

	producer, err := kafka.NewProducer("kafka:" + conf.KafkaBroker)
	if err != nil {
		logger.Error("unable to create a kafka producer: %s", err.Error())
	}
	defer producer.Close()

	useCase := usecase.NewParserUseCase(producer, conf.WhiteList)
	handler := handler.NewParserHandler(useCase)

	http.Handle("/", handler)
	logger.Error("%s", http.ListenAndServe(":"+conf.ParserPort, nil))
}
