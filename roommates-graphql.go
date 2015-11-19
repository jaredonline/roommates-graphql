package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// internal
	"github.com/jaredonline/roommates-graphql/data"
	"github.com/jaredonline/roommates-graphql/types"

	// external
	"github.com/graphql-go/graphql"
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

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()["query"][0]
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8888")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8888/graphql?query={users{first_name}}'")
	http.ListenAndServe(":8888", nil)
}
