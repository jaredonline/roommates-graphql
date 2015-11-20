package types

import (
	// external
	"github.com/graphql-go/graphql"
)

var Person = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
