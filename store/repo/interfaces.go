package repo

import "github.com/rls/ping-api/store/model"

// LocationSaver ..
type LocationSaver interface {
	Save(key string, locations ...*model.Location) error
	Get(key string, userIDs []string) (map[string]*model.Location, error)
}

// ILocation ...
type ILocation interface {
	LocationSaver
}
