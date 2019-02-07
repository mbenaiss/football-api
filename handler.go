package functions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-graphql-football-api/graphql"
)

func Matches(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("Method not allowed")
		w.Write([]byte(err.Error()))
		return
	}
	var params struct {
		Query string `json:"query"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(params.Query)
	response := graphql.Execute(params.Query)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
