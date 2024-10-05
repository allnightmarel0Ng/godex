package handler

import (
	"time"

	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/godex/internal/logger"
	confluentkafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ContainerKafkaHandler struct {
	consumer *kafka.Consumer
	useCase  usecase.ContainerUseCase
}

func NewContainerKafkaHandler(consumer *kafka.Consumer, useCase usecase.ContainerUseCase) ContainerKafkaHandler {
	return ContainerKafkaHandler{
		consumer: consumer,
		useCase:  useCase,
	}
}

func (c *ContainerKafkaHandler) Handle() {
	logger.Debug("Handle: start")
	defer logger.Debug("Handle: end")

	for {
		msg, err := c.consumer.Consume(time.Second)
		if err == nil {
			functionInfo := string(msg.Value)
			logger.Trace("got new function: %s", functionInfo)
			err = c.useCase.ProcessNewFunction(functionInfo)
			if err != nil {
				logger.Warning("got an error while storing the new function: %s", err.Error())
			} else {
				logger.Trace("new function was inserted successfully")
			}
		} else if !err.(confluentkafka.Error).IsTimeout() {
			logger.Warning("consumer error: %s", err.Error())
		}
	}
}
