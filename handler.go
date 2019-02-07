package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-graphql-football-api/graphql"
)

type Query struct {
	Query string `json:query`
}

func Matches(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		err := fmt.Errorf("Method not allowed")
		w.Write([]byte(err.Error()))
		return
	}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	query := &Query{}
	err = json.Unmarshal(b, query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	result := graphql.Execute(query.Query)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
