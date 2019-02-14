package queries

import (
	"strconv"
	"testing"

	"github.com/mbenaiss/football-api/models"
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

	testScenarios := []struct {
		given, expected []models.MatchResult
		f               map[string]interface{}
		expectedError   bool
	}{
		{
			given: matches,
			expected: []models.MatchResult{
				models.MatchResult{
					Competition: "Primera Division",
					Status:      "FINISHED",
				},
				models.MatchResult{
					Competition: "Primera Division",
					Status:      "SCHEDULED",
				}},
			expectedError: false,
			f: map[string]interface{}{
				"competition": "Primera Division",
			},
		},
		{
			f: map[string]interface{}{
				"competition": "Primera Division",
				"status":      "SCHEDULED",
			},
			expected: []models.MatchResult{
				models.MatchResult{
					Competition: "Primera Division",
					Status:      "SCHEDULED",
				},
			},
			expectedError: false,
		},
		{
			given:         matches,
			expected:      matches,
			f:             map[string]interface{}{},
			expectedError: false,
		},
		{
			given:         matches,
			expected:      matches,
			f:             map[string]interface{}{},
			expectedError: true,
		},
	}

	for i, s := range testScenarios {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			results := filter(matches, s.f)
			if len(s.expected) != len(results) && !s.expectedError {
				t.Errorf("expected len : %d \n %+v \ngot len : %d %+v", len(s.expected), s.expected, len(results), results)
			}
		})
	}
}
