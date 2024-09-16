package handler

import (
	"log"
	"time"

	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
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
	for {
		msg, err := c.consumer.Consume(time.Second)
		if err == nil {
			functionInfo := string(msg.Value)
			log.Printf("got new function: %s", functionInfo)
			err = c.useCase.ProcessNewFunction(functionInfo)
			if err != nil {
				log.Printf("got an error while storing the new function: %s", err.Error())
			}
		} else if !err.(confluentkafka.Error).IsTimeout() {
			log.Printf("consumer error: %s", err.Error())
		}
	}
}
