package queries

import (
	"github.com/go-graphql-football-api/data"
	"github.com/go-graphql-football-api/graphql/types"
	"github.com/graphql-go/graphql"
)

func GetMatchQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.MatchType),
		Args: graphql.FieldConfigArgument{
			"competition": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"status": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			matches, _ := data.GetDataFromApi()
			results := filter(matches, params.Args)
			return results, nil
		},
	}
}
