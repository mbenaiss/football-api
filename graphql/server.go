package graphql

import (
	"fmt"

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

func Execute(query string) *graphql.Result {
	schema, err := getSchema()
	if err != nil {
		panic(err)
	}
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v\n", result.Errors)
	}
	return result
}
