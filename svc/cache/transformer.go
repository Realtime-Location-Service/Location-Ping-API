package cache

import (
	"github.com/go-redis/redis"
	"github.com/rls/ping-api/store/model"
)

func transformToGeoLocation(locations ...*model.Location) []*redis.GeoLocation {
	geoLocs := []*redis.GeoLocation{}
	for _, loc := range locations {
		geoLocs = append(geoLocs,
			&redis.GeoLocation{Longitude: loc.Lon,
				Latitude: loc.Lat, Name: loc.UserID})
	}
	return geoLocs
}

func transformToUserLocation(userIDs []string, geoPos []*redis.GeoPos) map[string]*model.Location {
	locations := map[string]*model.Location{}
	for i, pos := range geoPos {
		if pos == nil {
			locations[userIDs[i]] = nil
			continue
		}
		locations[userIDs[i]] = &model.Location{Lat: pos.Latitude, Lon: pos.Longitude}
	}
	return locations
}

func transform(locs []redis.GeoLocation) []*model.Location {
	locations := []*model.Location{}
	for _, l := range locs {
		locations = append(locations, &model.Location{Lat: l.Latitude, Lon: l.Longitude,
			Dist: l.Dist, UserID: l.Name, GeoHash: l.GeoHash})
	}
	return locations
}
