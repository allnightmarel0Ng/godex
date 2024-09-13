package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/allnightmarel0Ng/godex/internal/app/parser/usecase"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ParserHandler struct {
	consumer *kafka.Consumer
	useCase  usecase.ParserUseCase
}

func NewParserHandler(consumer *kafka.Consumer, useCase usecase.ParserUseCase) ParserHandler {
	return ParserHandler{
		consumer: consumer,
		useCase:  useCase,
	}
}

func (p *ParserHandler) fetchFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", response.Status)
	}

	return io.ReadAll(response.Body)
}

func (p *ParserHandler) HandleMessage() {
	for {
		msg, err := p.consumer.ReadMessage(time.Second)
		if err == nil {
			url := msg.String()
		} else if !err.(*kafka.Error).IsTimeout() {
			log.Printf("Consumer error: %s", err.Error())
		}
	}
}
