package models

type Match struct {
	Count   int64          `json:"count"`
	Matches []MatchElement `json:"matches"`
}

type MatchElement struct {
	ID          int64         `json:"id"`
	Competition AwayTeamClass `json:"competition"`
	Status      Status        `json:"status"`
	Matchday    int64         `json:"matchday"`
	Stage       string        `json:"stage"`
	Group       string        `json:"group"`
	LastUpdated string        `json:"lastUpdated"`
	Score       Score         `json:"score"`
	HomeTeam    AwayTeamClass `json:"homeTeam"`
	AwayTeam    AwayTeamClass `json:"awayTeam"`
}

type AwayTeamClass struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Score struct {
	Winner    Winner    `json:"winner"`
	Duration  string    `json:"duration"`
	FullTime  ExtraTime `json:"fullTime"`
	HalfTime  ExtraTime `json:"halfTime"`
	ExtraTime ExtraTime `json:"extraTime"`
	Penalties ExtraTime `json:"penalties"`
}

type ExtraTime struct {
	HomeTeam int `json:"homeTeam"`
	AwayTeam int `json:"awayTeam"`
}

type Season struct {
	ID              int64       `json:"id"`
	StartDate       string      `json:"startDate"`
	EndDate         string      `json:"endDate"`
	CurrentMatchday int64       `json:"currentMatchday"`
	Winner          interface{} `json:"winner"`
}

type Winner string

const (
	AwayTeam Winner = "AWAY_TEAM"
	Draw     Winner = "DRAW"
	HomeTeam Winner = "HOME_TEAM"
)

type Status string

const (
	Finished  Status = "FINISHED"
	Scheduled Status = "SCHEDULED"
)
