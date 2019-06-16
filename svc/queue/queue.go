package queue

import (
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/utils/consts"
)

// IQueueService ...
type IQueueService interface {
	Publish(*model.Queue) error
}

// NewQueueService ...
func NewQueueService(qType consts.QueueType) IQueueService {
	if qType == consts.RabbitMQ {
		return NewRabbitMQ()
	}
	return nil
}
