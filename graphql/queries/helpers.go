package queries

import (
	"github.com/go-graphql-football-api/models"
)

func filter(matches []models.MatchResult, filters map[string]interface{}) []models.MatchResult {
	if len(filters) == 0 {
		return matches
	}
	results := []models.MatchResult{}

	for _, m := range matches {
		c := true
		s := true
		for k, v := range filters {
			if k == "competition" {
				if m.Competition != v.(string) {
					c = false
				}
			}
			if k == "status" {
				if m.Status != v.(string) {
					s = false
				}
			}
		}
		if c && s {
			results = append(results, m)
		}
	}
	return results
}
