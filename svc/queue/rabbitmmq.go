package queue

import (
	"github.com/rls/ping-api/conn/queue"
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/store/model"
	"github.com/streadway/amqp"
)

type channel struct {
	*amqp.Channel
	error
}

// RabbitMQ ...
type RabbitMQ struct {
	rmq *queue.RabbitMQ
	ch  *channel
}

func (r *RabbitMQ) getChannel() *channel {
	if r.ch.error != nil {
		c, err := r.rmq.Conn.Channel()
		r.ch = &channel{c, err}
	}
	return r.ch
}

// Publish sends message to queue
func (r *RabbitMQ) Publish(qm *model.Queue) error {

	ch := r.getChannel()
	if ch.error != nil {
		return ch.error
	}

	q, err := ch.QueueDeclare(
		qm.Name,    // name
		qm.Durable, // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		amqp.Table{
			"x-message-ttl": qm.TTL,
		},
	)

	if err != nil {
		return err
	}

	p := amqp.Publishing{
		ContentType: qm.ContentType,
		Body:        []byte(qm.Data),
	}

	if qm.Durable {
		p.DeliveryMode = amqp.Persistent
	}

	return ch.Publish(
		qm.Exchange, // exchange
		q.Name,      // routing key
		false,       // mandatory
		false,
		p)
}

// NewRabbitMQ returns rabbitmq connector
func NewRabbitMQ() IQueueService {
	c := queue.GetConnection(config.AppCfg().QueueType).(*queue.RabbitMQ)
	ch, err := c.Conn.Channel()
	return &RabbitMQ{c, &channel{ch, err}}
}
