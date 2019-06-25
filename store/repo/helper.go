package repo

import (
	"time"

	"strings"

	"github.com/rls/ping-api/store/model"
)

func (l *Location) resolveRequiredInfo(domain string, locations []*model.Location) []*model.Location {
	for _, l := range locations {
		l.ServerTimestampUTC = time.Now().UTC().Unix()
		l.Domain = domain
	}
	return locations
}

func (l *Location) getUniqueIDs(ids []string) []string {
	um := make(map[string]bool)
	uuIds := []string{}
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if _, ok := um[id]; !ok {
			uuIds = append(uuIds, id)
		}
		um[id] = true
	}
	return uuIds
}
