package cache

import (
	"github.com/rls/ping-api/conn/cache"
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/utils/errors"
)

// Redis ...
type Redis struct {
	redis *cache.Redis
}

// Get ...
func (r *Redis) Get(key string) (*model.Location, error) {
	return nil, nil
}

// GeoAdd adds locations in redis
// in case of error returns error with stack trace
func (r *Redis) GeoAdd(key string, locations ...*model.Location) error {
	if cmd := r.redis.GeoAdd(key, transformToGeoLocation(locations...)...); cmd.Err() != nil {
		return errors.Wrap(cmd.Err(), "error happend while adding geo location to redis")
	}
	return nil
}

// NewRedis returns redis client
func NewRedis() ICacheService {
	return &Redis{cache.GetClient(config.AppCfg().CacheType).(*cache.Redis)}
}
