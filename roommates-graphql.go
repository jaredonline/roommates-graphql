package main

import (
	"encoding/json"
	"fmt"
	"log"

	// internal
	"github.com/jaredonline/roommates-graphql/data"
	"github.com/jaredonline/roommates-graphql/types"

	// external
	"github.com/graphql-go/graphql"
	"gopkg.in/gorp.v1"
)

var (
	dbMap *gorp.DbMap
)

func main() {
	dbMap, err := data.InitDB("development")

	fields := graphql.Fields{
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

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// Query
	query := `
		query GetUsers {
			users {
				id
				first_name
				last_name
				age
			}
			jared: user(id:1) { first_name }
			cailin: user(id:3) { age }
		}
	`

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}

	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute GraphQL operaion: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)

	fmt.Println("%s", string(rJSON))
}
