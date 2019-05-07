package cache

import (
	"github.com/go-redis/redis"
	"github.com/rls/ping-api/pkg/config"
)

// Redis holds the configs
type Redis struct{ *redis.Client }

var redisDB = &Redis{}

// connects to redis
func (db *Redis) connect(cfg *redis.Options) error {
	db.Client = redis.NewClient(cfg)
	return nil
}

// GetRedis returns redis client
func GetRedis() *Redis {
	return redisDB
}

// ConnectRedis using cofigs
func ConnectRedis() error {
	return redisDB.connect(&config.RedisCfg().Opts)
}
