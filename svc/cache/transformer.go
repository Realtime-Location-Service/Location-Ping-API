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
