package repo

import "github.com/rls/ping-api/store/model"
import "github.com/rls/ping-api/svc/cache"

// Location ...
type Location struct {
	cacheSvc cache.ICacheService
}

// Save adds geo locations in cache
func (l *Location) Save(key string, locations ...*model.Location) error {
	return l.cacheSvc.GeoAdd(key, locations...)
}

// NewLocation returns a new location repo
func NewLocation(cacheSvc cache.ICacheService) ILocation {
	return &Location{cacheSvc}
}
