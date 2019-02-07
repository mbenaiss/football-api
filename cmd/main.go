package main

import (
	"log"
	"net/http"

	"github.com/go-graphql-football-api"
)

func main() {
	http.HandleFunc("/graphql", functions.Matches)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
