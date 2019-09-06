package cache

import (
	"github.com/go-redis/redis"
	"github.com/rls/ping-api/conn/cache"
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/utils/errors"
	"github.com/rls/ping-api/utils/hash"
	"strconv"
)

// Redis ...
type Redis struct {
	redis *cache.Redis
}

// Get returns locations of the users
func (r *Redis) Get(key string, userIDs ...string) (map[string]*model.Location, error) {
	geoPos, err := r.redis.GeoPos(key, userIDs...).Result()
	if err != nil {
		return nil, errors.Wrap(err, errors.ErrGettingGeoLocation)
	}
	return r.transformToUserLocation(key, userIDs, geoPos), nil
}

// GeoAdd adds locations in redis
// in case of error returns error with stack trace
func (r *Redis) GeoAdd(key string, locations ...*model.Location) error {
	if cmd := r.redis.GeoAdd(key, r.transformToGeoLocation(locations...)...); cmd.Err() != nil {
		return errors.Wrap(cmd.Err(), errors.ErrSavingGeoLocation)
	}
	return nil
}

// SaveLocationTimestamp keeps track of latest
// location's client timestamp
func (r *Redis) SaveLocationTimestamp(key string, location *model.Location) error {
	if cmd := r.redis.HSet(hash.GetMD5Hash(key), location.UserID, location.ClientTimestampUTC); cmd.Err() != nil {
		return errors.Wrap(cmd.Err(), errors.ErrSavingGeoLocationTimestamp)
	}
	return nil
}

// Search returns relevant  locations within radius
func (r *Redis) Search(key string, radius *model.Radius) ([]*model.Location, error) {
	res, err := r.redis.GeoRadius(key, radius.Lon, radius.Lat, &redis.GeoRadiusQuery{
		Radius:      radius.Val,
		Unit:        radius.Unit,
		WithGeoHash: true,
		WithCoord:   true,
		WithDist:    true,
		Sort:        "ASC",
		Count:       radius.Limit,
	}).Result()

	if err != nil {
		return nil, errors.Wrap(err, errors.ErrGeoRadiusSearch)
	}
	return r.transform(res), nil
}

// get timestamp of location
func (r *Redis) getLocationTimestamp(key, userID string) int64 {
	timestamp := r.redis.HGet(hash.GetMD5Hash(key), userID).Val()
	ts, _ := strconv.ParseInt(timestamp, 10, 64)
	return ts
}

// NewRedis returns redis client
func NewRedis() ICacheService {
	return &Redis{cache.GetClient(config.AppCfg().CacheType).(*cache.Redis)}
}
