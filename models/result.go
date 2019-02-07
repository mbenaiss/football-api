package models

type MatchResult struct {
	ID          int64  `json:"id"`
	Competition string `json:"competition"`
	Status      string `json:"status"`
	Matchday    int64  `json:"matchday"`
	Stage       string `json:"stage"`
	Group       string `json:"group"`
	LastUpdated string `json:"lastUpdated"`
	Score       string `json:"score"`
	HomeTeam    string `json:"homeTeam"`
	AwayTeam    string `json:"awayTeam"`
}
