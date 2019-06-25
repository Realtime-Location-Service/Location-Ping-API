package repo

import (
	"time"

	"github.com/rls/ping-api/store/model"
)

func (l *Location) resolveRequiredInfo(domain string, locations []*model.Location) []*model.Location {
	for _, l := range locations {
		l.ServerTimestampUTC = time.Now().UTC().Unix()
		l.Domain = domain
	}
	return locations
}
