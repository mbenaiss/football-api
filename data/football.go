package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-graphql-football-api/models"
)

func GetDataFromApi() ([]models.MatchResult, error) {
	url := "https://api.football-data.org/v2/teams/86/matches"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-auth-token", os.Getenv("FOOTBALL_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var match models.Match
	if err := json.NewDecoder(res.Body).Decode(&match); err != nil {
		return nil, err
	}
	return convertModel(match), nil
}

func convertModel(match models.Match) []models.MatchResult {
	results := []models.MatchResult{}
	for _, m := range match.Matches {
		n := models.MatchResult{
			ID:          m.ID,
			Matchday:    m.Matchday,
			AwayTeam:    m.AwayTeam.Name,
			HomeTeam:    m.HomeTeam.Name,
			Score:       fmt.Sprintf("%d : %d", m.Score.FullTime.HomeTeam, m.Score.FullTime.AwayTeam),
			Competition: m.Competition.Name,
			LastUpdated: m.LastUpdated,
			Group:       string(m.Group),
			Status:      string(m.Status),
			Stage:       string(m.Stage),
		}
		results = append(results, n)
	}
	return results
}
