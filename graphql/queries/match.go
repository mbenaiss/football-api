package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/mbenaiss/football-api/data"
	"github.com/mbenaiss/football-api/graphql/types"
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
