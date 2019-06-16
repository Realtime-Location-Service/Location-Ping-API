package queue

import (
	"github.com/pkg/errors"
	"github.com/rls/ping-api/utils/consts"
)

// Connect ...
func Connect(qt consts.QueueType) error {
	if qt == consts.RabbitMQ {
		return ConnectRabbitMQ()
	}
	return ErrInvalidQueueType
}

// GetConnection ...
func GetConnection(qt consts.QueueType) interface{} {
	if qt == consts.RabbitMQ {
		return GetRabbitMQ()
	}
	return ErrInvalidQueueType
}

var (
	// ErrInvalidQueueType ...
	ErrInvalidQueueType = errors.New("Invalid queue type")
)
