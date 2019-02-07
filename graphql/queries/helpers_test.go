package queries

import (
	"testing"

	"github.com/go-graphql-football-api/models"
)

func TestFilters(t *testing.T) {
	matches := []models.MatchResult{
		models.MatchResult{
			Competition: "CL",
			Status:      "FINISHED",
		},
		models.MatchResult{
			Competition: "CL",
			Status:      "FINISHED",
		},
		models.MatchResult{
			Competition: "CL",
			Status:      "SCHEDULED",
		},
		models.MatchResult{
			Competition: "Primera Division",
			Status:      "FINISHED",
		},
		models.MatchResult{
			Competition: "Primera Division",
			Status:      "SCHEDULED",
		},
	}

	results := filter(matches, map[string]interface{}{
		"competition": "Primera Division",
	})

	if len(results) < 2 {
		t.Errorf("expected 2 , got %+v", results)
	}

	results = filter(matches, map[string]interface{}{
		"competition": "Primera Division",
		"status":      "SCHEDULED",
	})

	if len(results) != 1 {
		t.Errorf("expected 1 , got len(%+v) %+v", len(results), results)
	}

	results = filter(matches, map[string]interface{}{})
	if len(results) != 5 {
		t.Errorf("expected 0 , got len(%+v) %+v", len(results), results)
	}
}
