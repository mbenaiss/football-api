package main

import (
	"log"
	"net/http"

	"github.com/mbenaiss/football-api/gcp"
)

func main() {
	http.HandleFunc("/graphql", gcp.Matches)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
