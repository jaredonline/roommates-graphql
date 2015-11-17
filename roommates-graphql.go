package main

import (
	"encoding/json"
	"fmt"
	"log"

	// external
	"github.com/graphql-go/graphql"
)

func main() {
	fields := graphql.FieldConfigMap{
		"hello": &graphql.FieldConfig{
			Type: graphql.String,
			Resolve: func(p graphql.GQLFRParams) interface{} {
				return "world"
			},
		},
		"foo": &graphql.FieldConfig{
			Type: graphql.String,
			Resolve: func(p graphql.GQLFRParams) interface{} {
				return "bar"
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// Query
	query := `
		{
			hello
			foo
		}
	`

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}

	r := graphql.Graphql(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute GraphQL operaion: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)

	fmt.Println("%s", string(rJSON))
}
