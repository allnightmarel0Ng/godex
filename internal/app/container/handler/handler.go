package handler

import (
	"github.com/allnightmarel0Ng/godex/internal/app/container/usecase"
	"github.com/allnightmarel0Ng/godex/internal/infrastructure/kafka"
	"github.com/allnightmarel0Ng/godex/internal/logger"
)

type ContainerHandler struct {
	consumer *kafka.Consumer
	useCase  usecase.ContainerUseCase
}

func NewContainerHandler(consumer *kafka.Consumer, useCase usecase.ContainerUseCase) ContainerHandler {
	return ContainerHandler{
		consumer: consumer,
		useCase:  useCase,
	}
}

func (c *ContainerHandler) Handle() {
	logger.Debug("Handle: start")
	defer logger.Debug("Handle: end")

	c.consumer.ConsumeMessagesEternally(c.useCase.ProcessNewFunction, logger.Trace, logger.Warning)
}
