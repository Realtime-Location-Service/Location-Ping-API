package repo

import (
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/svc/cache"
)

// Location ...
type Location struct {
	cacheSvc cache.ICacheService
}

// Save adds geo locations in cache
func (l *Location) Save(key string, locations ...*model.Location) error {
	return l.cacheSvc.GeoAdd(key, locations...)
}

// Get returns users locations
func (l *Location) Get(key string, userIDs []string) (map[string]*model.Location, error) {
	return l.cacheSvc.Get(key, getUniqueIDs(userIDs)...)
}

// NewLocation returns a new location repo
func NewLocation(cacheSvc cache.ICacheService) ILocation {
	return &Location{cacheSvc}
}
