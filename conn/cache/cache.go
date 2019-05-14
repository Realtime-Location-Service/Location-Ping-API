package cache

import (
	"github.com/pkg/errors"
	"github.com/rls/ping-api/utils/consts"
)

// Connect ...
func Connect(ct consts.CacheType) error {
	if ct == consts.Redis {
		return ConnectRedis()
	}
	return ErrInvalidCacheType
}

// GetClient ...
func GetClient(ct consts.CacheType) interface{} {
	if ct == consts.Redis {
		return GetRedis()
	}
	return ErrInvalidCacheType
}

var (
	// ErrInvalidCacheType ...
	ErrInvalidCacheType = errors.New("Invalid cache type")
)
