package main

import (
	"net/http"

	"github.com/allnightmarel0Ng/godex/internal/app/parser/handler"
	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
	"github.com/allnightmarel0Ng/godex/internal/config"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	router := gin.Default()

	router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost"},
        AllowMethods: []string{"POST"},
        AllowHeaders: []string{"Content-Type"},
    }))

	router.POST("/", handler.HandleLink)

	err = http.ListenAndServe(":"+conf.ParserPort, router)
	if err != nil {
		logger.Error("unable to listen and serve: %s", err.Error())
	}
}
