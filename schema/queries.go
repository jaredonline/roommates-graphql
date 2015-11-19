package schema

import (
	"log"

	// internal
	"github.com/jaredonline/roommates-graphql/data"
	"github.com/jaredonline/roommates-graphql/types"

	// external
	"github.com/graphql-go/graphql"
	"gopkg.in/gorp.v1"
)

var (
	dbMap  *gorp.DbMap
	fields graphql.Fields
	schema graphql.Schema
	err    error
)

func init() {
	dbMap, err = data.InitDB("development")
	fields = graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(types.User),
			Resolve: func(p graphql.ResolveParams) interface{} {
				return data.GetAllUsers(dbMap)
			},
		},
		"user": &graphql.Field{
			Type: types.User,
			Resolve: func(p graphql.ResolveParams) interface{} {
				if idQuery, isOK := p.Args["id"].(int); isOK {
					return data.GetUserById(dbMap, idQuery)
				}
				return nil
			},
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
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

	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}
}

func ExecuteQuery(query string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	return result
}
