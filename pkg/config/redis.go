package config

import (
	"time"

	redisc "github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// Redis holds the redis configuration
// https://godoc.org/github.com/go-redis/redis#Options
type Redis struct {
	Opts redisc.Options
}

var redis = &Redis{}

// RedisCfg returns the redis configuration
func RedisCfg() *Redis {
	return redis
}

// LoadRedisCfg loads redis configuration
func LoadRedisCfg() {
	redis.Opts.Network = viper.GetString("redis.network")
	redis.Opts.Addr = viper.GetString("redis.addr")
	redis.Opts.Password = viper.GetString("redis.password")
	redis.Opts.MaxRetries = viper.GetInt("redis.max_retries")
	redis.Opts.PoolSize = viper.GetInt("redis.pool_size")
	redis.Opts.MinIdleConns = viper.GetInt("redis.min_idle_conns")
	redis.Opts.MaxRetries = viper.GetInt("redis.max_retries")
	redis.Opts.MaxRetries = viper.GetInt("redis.max_retries")
	redis.Opts.ReadTimeout = viper.GetDuration("redis.read_timeout") * time.Second
	redis.Opts.WriteTimeout = viper.GetDuration("redis.write_timeout") * time.Second
	redis.Opts.IdleTimeout = viper.GetDuration("redis.idle_timeout") * time.Second
	redis.Opts.DialTimeout = viper.GetDuration("redis.dial_timeout") * time.Second
	redis.Opts.PoolTimeout = viper.GetDuration("redis.pool_timeout") * time.Second
}
