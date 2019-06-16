package queue

import (
	"github.com/rls/ping-api/pkg/config"
	"github.com/streadway/amqp"
)

// RabbitMQ holds the configs
type RabbitMQ struct{ Conn *amqp.Connection }

var rabbitMQ = &RabbitMQ{}

// connects to redis
func (r *RabbitMQ) connect(cfg *config.RabbitMQ) error {
	conn, err := amqp.Dial(cfg.URI)
	if err != nil {
		return err
	}
	r.Conn = conn

	return nil
}

// GetRabbitMQ returns rabbitmq connector
func GetRabbitMQ() *RabbitMQ {
	return rabbitMQ
}

// ConnectRabbitMQ using cofigs
func ConnectRabbitMQ() error {
	return rabbitMQ.connect(config.RabbitMQCfg())
}
