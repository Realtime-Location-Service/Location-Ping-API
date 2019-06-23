package queue

import (
	"log"

	"github.com/rls/ping-api/pkg/config"
	"github.com/streadway/amqp"
)

// RabbitMQ holds the configs
type RabbitMQ struct{ Conn *amqp.Connection }

var rabbitMQ = &RabbitMQ{}

// connects to redis
func (r *RabbitMQ) connect(cfg *config.RabbitMQ) error {
	c := make(chan *amqp.Error)
	go func() {
		err, ok := <-c
		if !ok {
			// On normal shutdowns, the chan will be closed.
			// so nothing to do
			return
		}
		log.Println("reconnecting: " + err.Error())
		r.connect(cfg)
	}()

	conn, err := amqp.Dial(cfg.URI)
	if err != nil {
		return err
	}
	r.Conn = conn
	conn.NotifyClose(c)

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
