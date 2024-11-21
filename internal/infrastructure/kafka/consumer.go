package kafka

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
}

func NewConsumer(broker string, group string) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: c,
	}, nil
}

func (c *Consumer) SubscribeTopics(topics []string) error {
	return c.consumer.SubscribeTopics(topics, nil)
}

func (c *Consumer) Consume(timeout time.Duration) (*kafka.Message, error) {
	return c.consumer.ReadMessage(timeout)
}

func (c *Consumer) Close() error {
	return c.consumer.Close()
}

func (c *Consumer) ConsumeMessagesEternally(dataCallback func([]byte) error, successCallback func(string, ...interface{}), errorCallback func(string, ...interface{})) {
	for {
		msg, err := c.Consume(time.Second)
		if err == nil {
			err = dataCallback(msg.Value)
			if err != nil {
				go errorCallback("got an error while storing the new function: %s", err.Error())
			} else {
				go successCallback("new function was inserted successfully")
			}
		} else if !err.(kafka.Error).IsTimeout() {
			go errorCallback("consumer error: %s", err.Error())
		}
	}
}
