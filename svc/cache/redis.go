package cache

import (
	"fmt"

	"github.com/rls/ping-api/conn/cache"
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/store/model"
)

// Redis ...
type Redis struct {
	conn *cache.Redis
}

// Get ...
func (r *Redis) Get(key string) (*model.Location, error) {
	return nil, nil
}

// GeoAdd ...
// TODO: need to implement
func (r *Redis) GeoAdd(key string, locations ...*model.Location) error {
	pong, err := r.conn.Ping().Result()
	fmt.Println(pong, err)
	return nil
}

// NewRedis ...
func NewRedis() ICacheService {
	return &Redis{cache.GetClient(config.AppCfg().CacheType).(*cache.Redis)}
}
