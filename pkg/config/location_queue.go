package config

import (
	"github.com/rls/ping-api/utils/consts"

	"github.com/spf13/viper"
)

// Queue ...
type Queue struct {
	Name          string
	ContentType   string
	Durable       bool
	AutoAck       bool
	Exchange      string
	PrefetchCount int
	PrefetchSize  int
}

// LocationQ ...
type LocationQ struct {
	Queue
}

var locationQ = &LocationQ{}

// LocationQCfg retuns configs
func LocationQCfg() *LocationQ {
	return locationQ
}

// LoadLocationQCfg loads configs
func LoadLocationQCfg() {
	locationQ.Name = viper.GetString("queue.location.name")
	locationQ.Exchange = viper.GetString("queue.location.exchange")
	locationQ.ContentType = consts.JSONContent
	locationQ.Durable = viper.GetBool("queue.location.durable")
	locationQ.AutoAck = viper.GetBool("queue.location.auto_ack")
	locationQ.PrefetchCount = viper.GetInt("queue.location.prefetch_count")
	locationQ.PrefetchSize = viper.GetInt("queue.location.prefetch_size")
}
