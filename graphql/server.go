package graphql

import (
	"github.com/go-graphql-football-api/graphql/queries"
	"github.com/graphql-go/graphql"
)

func getSchema() (graphql.Schema, error) {
	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queries.GetRootFields(),
		},
	)
}

func Execute(query string) (*graphql.Result, error) {
	schema, err := getSchema()
	if err != nil {
		return nil, err
	}
	return graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	}), nil
}
