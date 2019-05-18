package cache

import (
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/utils/consts"
)

// ICacheService ...
type ICacheService interface {
	Get(key string, userIDs ...string) (map[string]*model.Location, error)
	GeoAdd(key string, locations ...*model.Location) error
	Search(key string, radius *model.Radius) ([]*model.Location, error)
}

// NewCacheService ...
func NewCacheService(cacheType consts.CacheType) ICacheService {
	if cacheType == consts.Redis {
		return NewRedis()
	}
	return nil
}
