package messaging

import (
	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn *amqp091.Connection
}

func NewConsumer(url string) (*Consumer, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		conn: conn,
	}, nil
}

func (c *Consumer) Close() {
	c.conn.Close()
}

func (c *Consumer) Consume(q string) (<-chan amqp091.Delivery, func(), error) {
	ch, err := c.conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	msgs, err := ch.Consume(
		q,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	return msgs, func() {
		ch.Close()
	}, err
}
