package repo

import (
	"encoding/json"
	"log"

	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/store/model"
	"github.com/rls/ping-api/svc/cache"
	"github.com/rls/ping-api/svc/queue"
)

// Location ...
type Location struct {
	cacheSvc cache.ICacheService
	queueSvc queue.IQueueService
}

// Save adds geo locations in cache
func (l *Location) Save(key string, locations []*model.Location) error {
	if err := l.cacheSvc.GeoAdd(key, locations...); err != nil {
		return err
	}

	locations = l.resolveRequiredInfo(key, locations)
	ll, err := json.Marshal(locations)
	if err == nil {
		c := config.LocationQCfg()
		// ignore queueing error
		if err := l.queueSvc.Publish(&model.Queue{
			Name:        c.Name,
			Data:        ll,
			ContentType: c.ContentType,
			Durable:     c.Durable,
			Exchange:    c.Exchange,
			TTL:         c.TTL,
		}); err != nil {
			log.Println("error happened while publishing to queue reason: ", err)
		}
	}
	return nil
}

// Get returns users locations
func (l *Location) Get(key string, userIDs []string) (map[string]*model.Location, error) {
	return l.cacheSvc.Get(key, getUniqueIDs(userIDs)...)
}

// Search returns users locations within radius
func (l *Location) Search(key string, radius *model.Radius) ([]*model.Location, error) {
	return l.cacheSvc.Search(key, radius)
}

// NewLocation returns a new location repo
func NewLocation(cacheSvc cache.ICacheService, queueSvc queue.IQueueService) ILocation {
	return &Location{cacheSvc, queueSvc}
}
