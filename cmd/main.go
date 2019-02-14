package main

import (
	"log"
	"net/http"

	"github.com/mbenaiss/football-api/function"
)

func main() {
	http.HandleFunc("/graphql", function.Matches)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
