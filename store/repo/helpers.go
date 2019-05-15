package repo

import "strings"

func getUniqueIDs(ids []string) []string {
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
