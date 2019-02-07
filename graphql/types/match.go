package types

import (
	"github.com/graphql-go/graphql"
)

// define custom GraphQL ObjectType `MatchType` for our Golang struct `Match`
var MatchType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Match",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"competition": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"matchday": &graphql.Field{
			Type: graphql.Int,
		},
		"stage": &graphql.Field{
			Type: graphql.String,
		},
		"group": &graphql.Field{
			Type: graphql.String,
		},
		"lastUpdated": &graphql.Field{
			Type: graphql.String,
		},
		"score": &graphql.Field{
			Type: graphql.String,
		},
		"HomeTeam": &graphql.Field{
			Type: graphql.String,
		},
		"AwayTeam": &graphql.Field{
			Type: graphql.String,
		},
	},
})
