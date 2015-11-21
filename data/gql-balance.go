package data

import (
	// external
	"github.com/graphql-go/graphql"
)

var (
	Balance *graphql.Object
)

func init() {
	Balance = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Balance",
			Fields: graphql.Fields{
				"person": &graphql.Field{
					Type: Person,
				},
				"foreign_person": &graphql.Field{
					Type: Person,
				},
				"amount": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
