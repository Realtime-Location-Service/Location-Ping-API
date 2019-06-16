package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// RabbitMQ ...
type RabbitMQ struct {
	URI string
}

var rabbitMQ = &RabbitMQ{}

// RabbitMQCfg retuns configs
func RabbitMQCfg() *RabbitMQ {
	return rabbitMQ
}

// LoadRbbitMQCfg loads configs
func LoadRbbitMQCfg() {
	rabbitMQ.URI = fmt.Sprintf("%s://%s:%s@%s:%d",
		viper.GetString("rabbitmq.protocol"),
		viper.GetString("rabbitmq.username"),
		viper.GetString("rabbitmq.password"),
		viper.GetString("rabbitmq.host"),
		viper.GetInt("rabbitmq.port"))
}
