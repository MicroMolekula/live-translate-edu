package rabbitmq

import (
	"app-consumer/internal/configs"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

var declaredQueues = make(map[string]*amqp.Queue)

// #TODO Переписать Говно
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ(cfg *configs.Config) (*RabbitMQ, error) {
	connect, err := amqp.Dial(getUrlConnect(cfg))
	if err != nil {
		return nil, err
	}
	channel, err := connect.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{
		conn:    connect,
		channel: channel,
	}, nil
}

func (rmq *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
	queue, ok := declaredQueues[queueName]
	if !ok {
		q, err := rmq.channel.QueueDeclare(queueName, false, false, false, false, nil)
		if err != nil {
			return nil, err
		}
		queue = &q
	}
	messages, err := rmq.channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (rmq *RabbitMQ) Close() {
	if rmq.conn != nil {
		_ = rmq.conn.Close()
	}
	if rmq.channel != nil {
		_ = rmq.channel.Close()
	}
}

func getUrlConnect(cfg *configs.Config) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.Username,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
}
